你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# BOSH VirtualBox CPI

The BOSH VirtualBox CPI allows BOSH to manage *multiple* VirtualBox VMs/disks.

- [Download VirtualBox 5+](https://www.virtualbox.org/wiki/Downloads)
- [Cloud properties](docs/cloud-props.md)
- [Configuring Host-Only network](docs/networks-host-only.md) instructions to set up private VirtualBox network
- [Configuring NAT network](docs/networks-nat-network.md) instructions to set up public VirtualBox network

See [bosh-deployment's BOSH Lite on VirtualBox](https://github.com/cloudfoundry/bosh-deployment/blob/master/docs/bosh-lite-on-vbox.md) or [Concourse deployment](https://github.com/cppforlife/concourse-deployment) for example usage.

## TODO

- Aggressive VM deletion

```
CPI 'delete_vm' method responded with error: CmdError{"type":"Bosh::Clouds::CloudError","message":"Deleting vm 'vm-8b33e9d9-525f-49a9-6e1e-b156194ca0fe': Determining controller name: Retried '30' times: Running command: 'VBoxManage showvminfo vm-8b33e9d9-525f-49a9-6e1e-b156194ca0fe --machinereadable', stdout: '', stderr: 'VBoxManage: error: Could not find a registered machine named 'vm-8b33e9d9-525f-49a9-6e1e-b156194ca0fe'\nVBoxManage: error: Details: code VBOX_E_OBJECT_NOT_FOUND (0x80bb0001), component VirtualBoxWrap, interface IVirtualBox, callee nsISupports\nVBoxManage: error: Context: \"FindMachine(Bstr(VMNameOrUuid).raw(), machine.asOutParam())\" at line 2781 of file VBoxManageInfo.cpp\n': exit status 1","ok_to_retry":false}
```
