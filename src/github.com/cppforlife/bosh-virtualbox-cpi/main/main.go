package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshuuid "github.com/cloudfoundry/bosh-utils/uuid"
	"github.com/cppforlife/bosh-cpi-go/rpc"

	"github.com/cppforlife/bosh-virtualbox-cpi/cpi"
)

var (
	configPathOpt = flag.String("configPath", "", "Path to configuration file")
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // todo MAC generation

	logger, fs, cmdRunner, uuidGen := basicDeps()
	defer logger.HandlePanic("Main")

	flag.Parse()

	config, err := NewConfigFromPath(*configPathOpt, fs)
	if err != nil {
		logger.Error("main", "Loading config %s", err.Error())
		os.Exit(1)
	}

	compressor := boshcmd.NewTarballCompressor(cmdRunner, fs)

	cpiFactory := cpi.NewFactory(
		fs, cmdRunner, uuidGen, compressor, cpi.FactoryOpts(config), logger)

	cli := rpc.NewFactory(logger).NewCLI(cpiFactory)

	err = cli.ServeOnce()
	if err != nil {
		logger.Error("main", "Serving once: %s", err)
		os.Exit(1)
	}
}

func basicDeps() (boshlog.Logger, boshsys.FileSystem, boshsys.CmdRunner, boshuuid.Generator) {
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr)
	fs := boshsys.NewOsFileSystem(logger)
	cmdRunner := boshsys.NewExecCmdRunner(logger)
	uuidGen := boshuuid.NewGenerator()
	return logger, fs, cmdRunner, uuidGen
}
