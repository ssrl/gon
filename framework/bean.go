package bean

var Registry map[string]func()interface{} = map[string]func()interface{}{}

func GetBean(name string) interface{} {
	f := Registry[name]
	result := f()
	return result
}

func Bean(name string, f func()interface{}) {
	Registry[name] = f
}

func Ref(name string) interface{} {
	f := Registry[name]
	result := f()
	return result
}