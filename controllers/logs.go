package controllers

import (
	"encoding/json"
	"errors"
	"jobs/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//  LogsController operations for Logs
type LogsController struct {
	beego.Controller
}

// URLMapping ...
func (c *LogsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Logs
// @Param	body		body 	models.Joblog	true		"body for Logs content"
// @Success 201 {int} models.Joblog
// @Failure 403 body is empty
// @router / [post]
func (c *LogsController) Post() {
	var v models.Joblog
	empid := c.GetString("empid")
	empidint, _ := strconv.ParseInt(empid, 10, 64)
	moldid := c.GetString("moldid")
	moldidint, _ := strconv.ParseInt(moldid, 10, 64)
	jobid := c.GetString("jobid")
	jobidint, _ := strconv.ParseInt(jobid, 10, 64)
	tstarted := c.GetString("timestarted")
	ts, _ := time.Parse("2006-01-02T15:04", tstarted)
	tcompleted := c.GetString("timecompleted")
	tc, _ := time.Parse("2006-01-02T15:04", tcompleted)
	note := c.GetString("note")
	v.Empid = empidint
	v.Moldid = moldidint
	v.Jobid = jobidint
	v.Timestarted = ts
	v.Timecompleted = tc
	v.Note = note
	//json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddLogs(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Logs by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Joblog
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LogsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetLogsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Logs
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Joblog
// @Failure 403
// @router / [get]
func (c *LogsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllLogs(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Logs
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Joblog	true		"body for Logs content"
// @Success 200 {object} models.Joblog
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LogsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Joblog{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateLogsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Logs
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LogsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteLogs(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

/*	Create ...
	@Desc Create logs
*/
func (c *LogsController) Create() {
	c.Layout = "layout.html"
	c.TplName = "logs/create.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Scripts"] = "logs/createjs.tpl"
	c.Data["Titleicon"] = "添加工时记录"
}


func (c *LogsController) SearchByMold() {
	moldsn := c.GetString("mold")
	r, err := models.GetAllByMold(moldsn)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()
}

func (c *LogsController) SearchByEmp() {
	emp := c.GetString("emp")
	r, err := models.GetAllByEmp(emp)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = r
	}
	c.ServeJSON()
}

func (c *LogsController) Show() {
        c.Layout = "layout.html"
        c.TplName = "logs/show.tpl"
        c.LayoutSections = make(map[string]string)
        c.LayoutSections["Scripts"] = "logs/showjs.tpl"
        c.Data["Titleicon"] = "报表"
}