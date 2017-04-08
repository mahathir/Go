package utilities

type IocContainer struct {
	typeList map[string]interface{}
}

func NewIocContainer() *IocContainer {
	return &IocContainer{typeList: make(map[string]interface{})}
}

func (ioc *IocContainer) Reg(key string, theType interface{}) {
	ioc.typeList[key] = theType
}

func (ioc *IocContainer) Retrieve(key string) interface{} {
	return ioc.typeList[key]
}
