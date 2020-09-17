package core

var Event *EventManager


type EventManager struct {
	Event_UserCreate chan interface{}
	Event_UserEdit chan interface{}
	Event_UserDelete chan interface{}
	Event_RoleCreate chan interface{}
	Event_RoleEdit chan interface{}
	Event_RoleDelete chan interface{}
	Event_PermissionCreate chan interface{}
	Event_PermissionEdit chan interface{}
	Event_PermissionDelete chan interface{}
}


func init() {
	// Initialization Event
	// 初始化Event
	Event = &EventManager{
		Event_UserCreate:       make(chan interface{}, 10),
		Event_UserEdit:         make(chan interface{}, 10),
		Event_UserDelete:       make(chan interface{}, 10),
		Event_RoleCreate:       make(chan interface{}, 10),
		Event_RoleEdit:         make(chan interface{}, 10),
		Event_RoleDelete:       make(chan interface{}, 10),
		Event_PermissionCreate: make(chan interface{}, 10),
		Event_PermissionEdit:   make(chan interface{}, 10),
		Event_PermissionDelete: make(chan interface{}, 10),
	}
}


type Event_CallbackFunc func(interface{})

func (this *EventManager) UserCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserCreate
		callback(i)
	}
}

func (this *EventManager) UserEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserEdit
		callback(i)
	}
}

func (this *EventManager) UserDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_UserDelete
		callback(i)
	}
}

func (this *EventManager) RoleCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleCreate
		callback(i)
	}
}

func (this *EventManager) RoleEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleEdit
		callback(i)
	}
}

func (this *EventManager) RoleDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_RoleDelete
		callback(i)
	}
}

func (this *EventManager) PermissionCreate (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionCreate
		callback(i)
	}
}

func (this *EventManager) PermissionEdit (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionEdit
		callback(i)
	}
}

func (this *EventManager) PermissionDelete (callback Event_CallbackFunc)  {
	for {
		i := <- this.Event_PermissionDelete
		callback(i)
	}
}



