// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	app0xdd1446 "github.com/bitwormhole/bpm/app"
	command0xf0f741 "github.com/bitwormhole/bpm/command"
	service0xa5f732 "github.com/bitwormhole/bpm/service"
	cli0xf30272 "github.com/bitwormhole/starter-cli/cli"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-app0xdd1446.MainLoop
	cominfobuilder.Next()
	cominfobuilder.ID("com0-app0xdd1446.MainLoop").Class("looper").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMainLoop{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-command0xf0f741.BpmAutoUpgrade
	cominfobuilder.Next()
	cominfobuilder.ID("com1-command0xf0f741.BpmAutoUpgrade").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmAutoUpgrade{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com2-command0xf0f741.BpmFetch
	cominfobuilder.Next()
	cominfobuilder.ID("com2-command0xf0f741.BpmFetch").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmFetch{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com3-command0xf0f741.BpmHelp
	cominfobuilder.Next()
	cominfobuilder.ID("com3-command0xf0f741.BpmHelp").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmHelp{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com4-command0xf0f741.BpmInstall
	cominfobuilder.Next()
	cominfobuilder.ID("com4-command0xf0f741.BpmInstall").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmInstall{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com5-command0xf0f741.BpmMake
	cominfobuilder.Next()
	cominfobuilder.ID("com5-command0xf0f741.BpmMake").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmMake{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com6-command0xf0f741.BpmPackInfo
	cominfobuilder.Next()
	cominfobuilder.ID("com6-command0xf0f741.BpmPackInfo").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmPackInfo{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com7-command0xf0f741.BpmRun
	cominfobuilder.Next()
	cominfobuilder.ID("com7-command0xf0f741.BpmRun").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmRun{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com8-command0xf0f741.BpmUpdate
	cominfobuilder.Next()
	cominfobuilder.ID("com8-command0xf0f741.BpmUpdate").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmUpdate{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com9-command0xf0f741.BpmUpgrade
	cominfobuilder.Next()
	cominfobuilder.ID("com9-command0xf0f741.BpmUpgrade").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmUpgrade{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com10-command0xf0f741.BpmVersion
	cominfobuilder.Next()
	cominfobuilder.ID("com10-command0xf0f741.BpmVersion").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpmVersion{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com11-command0xf0f741.Bpm
	cominfobuilder.Next()
	cominfobuilder.ID("com11-command0xf0f741.Bpm").Class("cli-handler").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComBpm{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-deploy-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-deploy-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComDeployServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-fetch-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-fetch-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFetchServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-install-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-install-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComInstallServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-make-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-make-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComMakeServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-pack-info-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-pack-info-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPackInfoServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-run-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-run-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComRunServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-update-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-update-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUpdateServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-upgrade-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-upgrade-service").Class("bpm-service").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComUpgradeServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-env-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-env-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComEnvServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-package-manager
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-package-manager").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComPackageManagerImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: bpm-remote-service
	cominfobuilder.Next()
	cominfobuilder.ID("bpm-remote-service").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHTTPGetServiceImpl{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMainLoop : the factory of component: com0-app0xdd1446.MainLoop
type comFactory4pComMainLoop struct {

    mPrototype * app0xdd1446.MainLoop

	
	mClientFactorySelector config.InjectionSelector
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComMainLoop) init() application.ComponentFactory {

	
	inst.mClientFactorySelector = config.NewInjectionSelector("#cli-client-factory",nil)
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMainLoop) newObject() * app0xdd1446.MainLoop {
	return & app0xdd1446.MainLoop {}
}

func (inst * comFactory4pComMainLoop) castObject(instance application.ComponentInstance) * app0xdd1446.MainLoop {
	return instance.Get().(*app0xdd1446.MainLoop)
}

func (inst * comFactory4pComMainLoop) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMainLoop) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMainLoop) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMainLoop) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainLoop) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMainLoop) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.ClientFactory = inst.getterForFieldClientFactorySelector(context)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldClientFactorySelector
func (inst * comFactory4pComMainLoop) getterForFieldClientFactorySelector (context application.InstanceContext) cli0xf30272.ClientFactory {

	o1 := inst.mClientFactorySelector.GetOne(context)
	o2, ok := o1.(cli0xf30272.ClientFactory)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-app0xdd1446.MainLoop")
		eb.Set("field", "ClientFactory")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf30272.ClientFactory")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldContextSelector
func (inst * comFactory4pComMainLoop) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmAutoUpgrade : the factory of component: com1-command0xf0f741.BpmAutoUpgrade
type comFactory4pComBpmAutoUpgrade struct {

    mPrototype * command0xf0f741.BpmAutoUpgrade

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmAutoUpgrade) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-upgrade-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmAutoUpgrade) newObject() * command0xf0f741.BpmAutoUpgrade {
	return & command0xf0f741.BpmAutoUpgrade {}
}

func (inst * comFactory4pComBpmAutoUpgrade) castObject(instance application.ComponentInstance) * command0xf0f741.BpmAutoUpgrade {
	return instance.Get().(*command0xf0f741.BpmAutoUpgrade)
}

func (inst * comFactory4pComBpmAutoUpgrade) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmAutoUpgrade) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmAutoUpgrade) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmAutoUpgrade) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmAutoUpgrade) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmAutoUpgrade) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmAutoUpgrade) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.UpgradeService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.UpgradeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com1-command0xf0f741.BpmAutoUpgrade")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.UpgradeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmFetch : the factory of component: com2-command0xf0f741.BpmFetch
type comFactory4pComBpmFetch struct {

    mPrototype * command0xf0f741.BpmFetch

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmFetch) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-fetch-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmFetch) newObject() * command0xf0f741.BpmFetch {
	return & command0xf0f741.BpmFetch {}
}

func (inst * comFactory4pComBpmFetch) castObject(instance application.ComponentInstance) * command0xf0f741.BpmFetch {
	return instance.Get().(*command0xf0f741.BpmFetch)
}

func (inst * comFactory4pComBpmFetch) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmFetch) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmFetch) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmFetch) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmFetch) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmFetch) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmFetch) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.FetchService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.FetchService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com2-command0xf0f741.BpmFetch")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.FetchService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmHelp : the factory of component: com3-command0xf0f741.BpmHelp
type comFactory4pComBpmHelp struct {

    mPrototype * command0xf0f741.BpmHelp

	

}

func (inst * comFactory4pComBpmHelp) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmHelp) newObject() * command0xf0f741.BpmHelp {
	return & command0xf0f741.BpmHelp {}
}

func (inst * comFactory4pComBpmHelp) castObject(instance application.ComponentInstance) * command0xf0f741.BpmHelp {
	return instance.Get().(*command0xf0f741.BpmHelp)
}

func (inst * comFactory4pComBpmHelp) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmHelp) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmHelp) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmHelp) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmHelp) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmHelp) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmInstall : the factory of component: com4-command0xf0f741.BpmInstall
type comFactory4pComBpmInstall struct {

    mPrototype * command0xf0f741.BpmInstall

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmInstall) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-install-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmInstall) newObject() * command0xf0f741.BpmInstall {
	return & command0xf0f741.BpmInstall {}
}

func (inst * comFactory4pComBpmInstall) castObject(instance application.ComponentInstance) * command0xf0f741.BpmInstall {
	return instance.Get().(*command0xf0f741.BpmInstall)
}

func (inst * comFactory4pComBpmInstall) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmInstall) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmInstall) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmInstall) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmInstall) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmInstall) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmInstall) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.InstallService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.InstallService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com4-command0xf0f741.BpmInstall")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.InstallService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmMake : the factory of component: com5-command0xf0f741.BpmMake
type comFactory4pComBpmMake struct {

    mPrototype * command0xf0f741.BpmMake

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmMake) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-make-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmMake) newObject() * command0xf0f741.BpmMake {
	return & command0xf0f741.BpmMake {}
}

func (inst * comFactory4pComBpmMake) castObject(instance application.ComponentInstance) * command0xf0f741.BpmMake {
	return instance.Get().(*command0xf0f741.BpmMake)
}

func (inst * comFactory4pComBpmMake) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmMake) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmMake) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmMake) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmMake) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmMake) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmMake) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.MakeService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.MakeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com5-command0xf0f741.BpmMake")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.MakeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmPackInfo : the factory of component: com6-command0xf0f741.BpmPackInfo
type comFactory4pComBpmPackInfo struct {

    mPrototype * command0xf0f741.BpmPackInfo

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmPackInfo) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-pack-info-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmPackInfo) newObject() * command0xf0f741.BpmPackInfo {
	return & command0xf0f741.BpmPackInfo {}
}

func (inst * comFactory4pComBpmPackInfo) castObject(instance application.ComponentInstance) * command0xf0f741.BpmPackInfo {
	return instance.Get().(*command0xf0f741.BpmPackInfo)
}

func (inst * comFactory4pComBpmPackInfo) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmPackInfo) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmPackInfo) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmPackInfo) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmPackInfo) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmPackInfo) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmPackInfo) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.PackInfoService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackInfoService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com6-command0xf0f741.BpmPackInfo")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackInfoService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmRun : the factory of component: com7-command0xf0f741.BpmRun
type comFactory4pComBpmRun struct {

    mPrototype * command0xf0f741.BpmRun

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmRun) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-run-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmRun) newObject() * command0xf0f741.BpmRun {
	return & command0xf0f741.BpmRun {}
}

func (inst * comFactory4pComBpmRun) castObject(instance application.ComponentInstance) * command0xf0f741.BpmRun {
	return instance.Get().(*command0xf0f741.BpmRun)
}

func (inst * comFactory4pComBpmRun) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmRun) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmRun) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmRun) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmRun) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmRun) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmRun) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.RunService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.RunService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com7-command0xf0f741.BpmRun")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.RunService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmUpdate : the factory of component: com8-command0xf0f741.BpmUpdate
type comFactory4pComBpmUpdate struct {

    mPrototype * command0xf0f741.BpmUpdate

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmUpdate) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-update-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmUpdate) newObject() * command0xf0f741.BpmUpdate {
	return & command0xf0f741.BpmUpdate {}
}

func (inst * comFactory4pComBpmUpdate) castObject(instance application.ComponentInstance) * command0xf0f741.BpmUpdate {
	return instance.Get().(*command0xf0f741.BpmUpdate)
}

func (inst * comFactory4pComBpmUpdate) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmUpdate) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmUpdate) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmUpdate) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmUpdate) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmUpdate) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmUpdate) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.UpdateService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.UpdateService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com8-command0xf0f741.BpmUpdate")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.UpdateService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmUpgrade : the factory of component: com9-command0xf0f741.BpmUpgrade
type comFactory4pComBpmUpgrade struct {

    mPrototype * command0xf0f741.BpmUpgrade

	
	mServiceSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmUpgrade) init() application.ComponentFactory {

	
	inst.mServiceSelector = config.NewInjectionSelector("#bpm-upgrade-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmUpgrade) newObject() * command0xf0f741.BpmUpgrade {
	return & command0xf0f741.BpmUpgrade {}
}

func (inst * comFactory4pComBpmUpgrade) castObject(instance application.ComponentInstance) * command0xf0f741.BpmUpgrade {
	return instance.Get().(*command0xf0f741.BpmUpgrade)
}

func (inst * comFactory4pComBpmUpgrade) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmUpgrade) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmUpgrade) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmUpgrade) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmUpgrade) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmUpgrade) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Service = inst.getterForFieldServiceSelector(context)
	return context.LastError()
}

//getterForFieldServiceSelector
func (inst * comFactory4pComBpmUpgrade) getterForFieldServiceSelector (context application.InstanceContext) service0xa5f732.UpgradeService {

	o1 := inst.mServiceSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.UpgradeService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com9-command0xf0f741.BpmUpgrade")
		eb.Set("field", "Service")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.UpgradeService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpmVersion : the factory of component: com10-command0xf0f741.BpmVersion
type comFactory4pComBpmVersion struct {

    mPrototype * command0xf0f741.BpmVersion

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComBpmVersion) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpmVersion) newObject() * command0xf0f741.BpmVersion {
	return & command0xf0f741.BpmVersion {}
}

func (inst * comFactory4pComBpmVersion) castObject(instance application.ComponentInstance) * command0xf0f741.BpmVersion {
	return instance.Get().(*command0xf0f741.BpmVersion)
}

func (inst * comFactory4pComBpmVersion) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpmVersion) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpmVersion) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpmVersion) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmVersion) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpmVersion) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComBpmVersion) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComBpm : the factory of component: com11-command0xf0f741.Bpm
type comFactory4pComBpm struct {

    mPrototype * command0xf0f741.Bpm

	

}

func (inst * comFactory4pComBpm) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComBpm) newObject() * command0xf0f741.Bpm {
	return & command0xf0f741.Bpm {}
}

func (inst * comFactory4pComBpm) castObject(instance application.ComponentInstance) * command0xf0f741.Bpm {
	return instance.Get().(*command0xf0f741.Bpm)
}

func (inst * comFactory4pComBpm) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComBpm) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComBpm) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComBpm) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpm) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComBpm) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComDeployServiceImpl : the factory of component: bpm-deploy-service
type comFactory4pComDeployServiceImpl struct {

    mPrototype * service0xa5f732.DeployServiceImpl

	
	mPMSelector config.InjectionSelector
	mEnvSelector config.InjectionSelector

}

func (inst * comFactory4pComDeployServiceImpl) init() application.ComponentFactory {

	
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)
	inst.mEnvSelector = config.NewInjectionSelector("#bpm-env-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComDeployServiceImpl) newObject() * service0xa5f732.DeployServiceImpl {
	return & service0xa5f732.DeployServiceImpl {}
}

func (inst * comFactory4pComDeployServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.DeployServiceImpl {
	return instance.Get().(*service0xa5f732.DeployServiceImpl)
}

func (inst * comFactory4pComDeployServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComDeployServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComDeployServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComDeployServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDeployServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComDeployServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PM = inst.getterForFieldPMSelector(context)
	obj.Env = inst.getterForFieldEnvSelector(context)
	return context.LastError()
}

//getterForFieldPMSelector
func (inst * comFactory4pComDeployServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-deploy-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnvSelector
func (inst * comFactory4pComDeployServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) service0xa5f732.EnvService {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.EnvService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-deploy-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.EnvService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFetchServiceImpl : the factory of component: bpm-fetch-service
type comFactory4pComFetchServiceImpl struct {

    mPrototype * service0xa5f732.FetchServiceImpl

	
	mPMSelector config.InjectionSelector
	mEnvSelector config.InjectionSelector
	mRemoteSelector config.InjectionSelector

}

func (inst * comFactory4pComFetchServiceImpl) init() application.ComponentFactory {

	
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)
	inst.mEnvSelector = config.NewInjectionSelector("#bpm-env-service",nil)
	inst.mRemoteSelector = config.NewInjectionSelector("#bpm-remote-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFetchServiceImpl) newObject() * service0xa5f732.FetchServiceImpl {
	return & service0xa5f732.FetchServiceImpl {}
}

func (inst * comFactory4pComFetchServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.FetchServiceImpl {
	return instance.Get().(*service0xa5f732.FetchServiceImpl)
}

func (inst * comFactory4pComFetchServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFetchServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFetchServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFetchServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFetchServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFetchServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PM = inst.getterForFieldPMSelector(context)
	obj.Env = inst.getterForFieldEnvSelector(context)
	obj.Remote = inst.getterForFieldRemoteSelector(context)
	return context.LastError()
}

//getterForFieldPMSelector
func (inst * comFactory4pComFetchServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-fetch-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnvSelector
func (inst * comFactory4pComFetchServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) service0xa5f732.EnvService {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.EnvService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-fetch-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.EnvService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldRemoteSelector
func (inst * comFactory4pComFetchServiceImpl) getterForFieldRemoteSelector (context application.InstanceContext) service0xa5f732.HTTPGetService {

	o1 := inst.mRemoteSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.HTTPGetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-fetch-service")
		eb.Set("field", "Remote")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.HTTPGetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComInstallServiceImpl : the factory of component: bpm-install-service
type comFactory4pComInstallServiceImpl struct {

    mPrototype * service0xa5f732.InstallServiceImpl

	
	mPMSelector config.InjectionSelector
	mFetchSerSelector config.InjectionSelector
	mDeploySerSelector config.InjectionSelector

}

func (inst * comFactory4pComInstallServiceImpl) init() application.ComponentFactory {

	
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)
	inst.mFetchSerSelector = config.NewInjectionSelector("#bpm-fetch-service",nil)
	inst.mDeploySerSelector = config.NewInjectionSelector("#bpm-deploy-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComInstallServiceImpl) newObject() * service0xa5f732.InstallServiceImpl {
	return & service0xa5f732.InstallServiceImpl {}
}

func (inst * comFactory4pComInstallServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.InstallServiceImpl {
	return instance.Get().(*service0xa5f732.InstallServiceImpl)
}

func (inst * comFactory4pComInstallServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComInstallServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComInstallServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComInstallServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComInstallServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComInstallServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PM = inst.getterForFieldPMSelector(context)
	obj.FetchSer = inst.getterForFieldFetchSerSelector(context)
	obj.DeploySer = inst.getterForFieldDeploySerSelector(context)
	return context.LastError()
}

//getterForFieldPMSelector
func (inst * comFactory4pComInstallServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-install-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFetchSerSelector
func (inst * comFactory4pComInstallServiceImpl) getterForFieldFetchSerSelector (context application.InstanceContext) service0xa5f732.FetchService {

	o1 := inst.mFetchSerSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.FetchService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-install-service")
		eb.Set("field", "FetchSer")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.FetchService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDeploySerSelector
func (inst * comFactory4pComInstallServiceImpl) getterForFieldDeploySerSelector (context application.InstanceContext) service0xa5f732.DeployService {

	o1 := inst.mDeploySerSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.DeployService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-install-service")
		eb.Set("field", "DeploySer")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.DeployService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComMakeServiceImpl : the factory of component: bpm-make-service
type comFactory4pComMakeServiceImpl struct {

    mPrototype * service0xa5f732.MakeServiceImpl

	

}

func (inst * comFactory4pComMakeServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComMakeServiceImpl) newObject() * service0xa5f732.MakeServiceImpl {
	return & service0xa5f732.MakeServiceImpl {}
}

func (inst * comFactory4pComMakeServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.MakeServiceImpl {
	return instance.Get().(*service0xa5f732.MakeServiceImpl)
}

func (inst * comFactory4pComMakeServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComMakeServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComMakeServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComMakeServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMakeServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComMakeServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPackInfoServiceImpl : the factory of component: bpm-pack-info-service
type comFactory4pComPackInfoServiceImpl struct {

    mPrototype * service0xa5f732.PackInfoServiceImpl

	
	mPMSelector config.InjectionSelector

}

func (inst * comFactory4pComPackInfoServiceImpl) init() application.ComponentFactory {

	
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPackInfoServiceImpl) newObject() * service0xa5f732.PackInfoServiceImpl {
	return & service0xa5f732.PackInfoServiceImpl {}
}

func (inst * comFactory4pComPackInfoServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.PackInfoServiceImpl {
	return instance.Get().(*service0xa5f732.PackInfoServiceImpl)
}

func (inst * comFactory4pComPackInfoServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPackInfoServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPackInfoServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPackInfoServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackInfoServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackInfoServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PM = inst.getterForFieldPMSelector(context)
	return context.LastError()
}

//getterForFieldPMSelector
func (inst * comFactory4pComPackInfoServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-pack-info-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComRunServiceImpl : the factory of component: bpm-run-service
type comFactory4pComRunServiceImpl struct {

    mPrototype * service0xa5f732.RunServiceImpl

	
	mPMSelector config.InjectionSelector
	mEnvSelector config.InjectionSelector

}

func (inst * comFactory4pComRunServiceImpl) init() application.ComponentFactory {

	
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)
	inst.mEnvSelector = config.NewInjectionSelector("#bpm-env-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComRunServiceImpl) newObject() * service0xa5f732.RunServiceImpl {
	return & service0xa5f732.RunServiceImpl {}
}

func (inst * comFactory4pComRunServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.RunServiceImpl {
	return instance.Get().(*service0xa5f732.RunServiceImpl)
}

func (inst * comFactory4pComRunServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComRunServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComRunServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComRunServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComRunServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.PM = inst.getterForFieldPMSelector(context)
	obj.Env = inst.getterForFieldEnvSelector(context)
	return context.LastError()
}

//getterForFieldPMSelector
func (inst * comFactory4pComRunServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-run-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldEnvSelector
func (inst * comFactory4pComRunServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) service0xa5f732.EnvService {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.EnvService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-run-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.EnvService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUpdateServiceImpl : the factory of component: bpm-update-service
type comFactory4pComUpdateServiceImpl struct {

    mPrototype * service0xa5f732.UpdateServiceImpl

	
	mRemoteSelector config.InjectionSelector
	mPMSelector config.InjectionSelector

}

func (inst * comFactory4pComUpdateServiceImpl) init() application.ComponentFactory {

	
	inst.mRemoteSelector = config.NewInjectionSelector("#bpm-remote-service",nil)
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUpdateServiceImpl) newObject() * service0xa5f732.UpdateServiceImpl {
	return & service0xa5f732.UpdateServiceImpl {}
}

func (inst * comFactory4pComUpdateServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.UpdateServiceImpl {
	return instance.Get().(*service0xa5f732.UpdateServiceImpl)
}

func (inst * comFactory4pComUpdateServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUpdateServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUpdateServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUpdateServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUpdateServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUpdateServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Remote = inst.getterForFieldRemoteSelector(context)
	obj.PM = inst.getterForFieldPMSelector(context)
	return context.LastError()
}

//getterForFieldRemoteSelector
func (inst * comFactory4pComUpdateServiceImpl) getterForFieldRemoteSelector (context application.InstanceContext) service0xa5f732.HTTPGetService {

	o1 := inst.mRemoteSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.HTTPGetService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-update-service")
		eb.Set("field", "Remote")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.HTTPGetService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPMSelector
func (inst * comFactory4pComUpdateServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-update-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComUpgradeServiceImpl : the factory of component: bpm-upgrade-service
type comFactory4pComUpgradeServiceImpl struct {

    mPrototype * service0xa5f732.UpgradeServiceImpl

	
	mEnvSelector config.InjectionSelector
	mPMSelector config.InjectionSelector
	mFetchSerSelector config.InjectionSelector
	mDeploySerSelector config.InjectionSelector

}

func (inst * comFactory4pComUpgradeServiceImpl) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#bpm-env-service",nil)
	inst.mPMSelector = config.NewInjectionSelector("#bpm-package-manager",nil)
	inst.mFetchSerSelector = config.NewInjectionSelector("#bpm-fetch-service",nil)
	inst.mDeploySerSelector = config.NewInjectionSelector("#bpm-deploy-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComUpgradeServiceImpl) newObject() * service0xa5f732.UpgradeServiceImpl {
	return & service0xa5f732.UpgradeServiceImpl {}
}

func (inst * comFactory4pComUpgradeServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.UpgradeServiceImpl {
	return instance.Get().(*service0xa5f732.UpgradeServiceImpl)
}

func (inst * comFactory4pComUpgradeServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComUpgradeServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComUpgradeServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComUpgradeServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUpgradeServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComUpgradeServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Env = inst.getterForFieldEnvSelector(context)
	obj.PM = inst.getterForFieldPMSelector(context)
	obj.FetchSer = inst.getterForFieldFetchSerSelector(context)
	obj.DeploySer = inst.getterForFieldDeploySerSelector(context)
	return context.LastError()
}

//getterForFieldEnvSelector
func (inst * comFactory4pComUpgradeServiceImpl) getterForFieldEnvSelector (context application.InstanceContext) service0xa5f732.EnvService {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.EnvService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-upgrade-service")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.EnvService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldPMSelector
func (inst * comFactory4pComUpgradeServiceImpl) getterForFieldPMSelector (context application.InstanceContext) service0xa5f732.PackageManager {

	o1 := inst.mPMSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.PackageManager)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-upgrade-service")
		eb.Set("field", "PM")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.PackageManager")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldFetchSerSelector
func (inst * comFactory4pComUpgradeServiceImpl) getterForFieldFetchSerSelector (context application.InstanceContext) service0xa5f732.FetchService {

	o1 := inst.mFetchSerSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.FetchService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-upgrade-service")
		eb.Set("field", "FetchSer")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.FetchService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}

//getterForFieldDeploySerSelector
func (inst * comFactory4pComUpgradeServiceImpl) getterForFieldDeploySerSelector (context application.InstanceContext) service0xa5f732.DeployService {

	o1 := inst.mDeploySerSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.DeployService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-upgrade-service")
		eb.Set("field", "DeploySer")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.DeployService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComEnvServiceImpl : the factory of component: bpm-env-service
type comFactory4pComEnvServiceImpl struct {

    mPrototype * service0xa5f732.EnvServiceImpl

	
	mContextSelector config.InjectionSelector

}

func (inst * comFactory4pComEnvServiceImpl) init() application.ComponentFactory {

	
	inst.mContextSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComEnvServiceImpl) newObject() * service0xa5f732.EnvServiceImpl {
	return & service0xa5f732.EnvServiceImpl {}
}

func (inst * comFactory4pComEnvServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.EnvServiceImpl {
	return instance.Get().(*service0xa5f732.EnvServiceImpl)
}

func (inst * comFactory4pComEnvServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComEnvServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComEnvServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComEnvServiceImpl) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst * comFactory4pComEnvServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComEnvServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Context = inst.getterForFieldContextSelector(context)
	return context.LastError()
}

//getterForFieldContextSelector
func (inst * comFactory4pComEnvServiceImpl) getterForFieldContextSelector (context application.InstanceContext) application.Context {
    return context.Context()
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComPackageManagerImpl : the factory of component: bpm-package-manager
type comFactory4pComPackageManagerImpl struct {

    mPrototype * service0xa5f732.PackageManagerImpl

	
	mEnvSelector config.InjectionSelector

}

func (inst * comFactory4pComPackageManagerImpl) init() application.ComponentFactory {

	
	inst.mEnvSelector = config.NewInjectionSelector("#bpm-env-service",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComPackageManagerImpl) newObject() * service0xa5f732.PackageManagerImpl {
	return & service0xa5f732.PackageManagerImpl {}
}

func (inst * comFactory4pComPackageManagerImpl) castObject(instance application.ComponentInstance) * service0xa5f732.PackageManagerImpl {
	return instance.Get().(*service0xa5f732.PackageManagerImpl)
}

func (inst * comFactory4pComPackageManagerImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComPackageManagerImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComPackageManagerImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComPackageManagerImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageManagerImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComPackageManagerImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Env = inst.getterForFieldEnvSelector(context)
	return context.LastError()
}

//getterForFieldEnvSelector
func (inst * comFactory4pComPackageManagerImpl) getterForFieldEnvSelector (context application.InstanceContext) service0xa5f732.EnvService {

	o1 := inst.mEnvSelector.GetOne(context)
	o2, ok := o1.(service0xa5f732.EnvService)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "bpm-package-manager")
		eb.Set("field", "Env")
		eb.Set("type1", "?")
		eb.Set("type2", "service0xa5f732.EnvService")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHTTPGetServiceImpl : the factory of component: bpm-remote-service
type comFactory4pComHTTPGetServiceImpl struct {

    mPrototype * service0xa5f732.HTTPGetServiceImpl

	

}

func (inst * comFactory4pComHTTPGetServiceImpl) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHTTPGetServiceImpl) newObject() * service0xa5f732.HTTPGetServiceImpl {
	return & service0xa5f732.HTTPGetServiceImpl {}
}

func (inst * comFactory4pComHTTPGetServiceImpl) castObject(instance application.ComponentInstance) * service0xa5f732.HTTPGetServiceImpl {
	return instance.Get().(*service0xa5f732.HTTPGetServiceImpl)
}

func (inst * comFactory4pComHTTPGetServiceImpl) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHTTPGetServiceImpl) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHTTPGetServiceImpl) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHTTPGetServiceImpl) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGetServiceImpl) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHTTPGetServiceImpl) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




