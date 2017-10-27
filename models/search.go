/* 
* @Author: 29575
* @Date:   2017-07-27 22:44:40
* @Last Modified by:   anchen
* @Last Modified time: 2017-08-11 14:13:40
*/

package models

import (
    "github.com/astaxie/beego/orm"
)

type Result struct{
    Empno   string
    Empname string
    Moldsn  string
    Moldname    string
    Codeno  string
    Descrip string
    Timestarted string
    Timecompleted  string
    Cal float64
}

func init() {

}

//通过模具编号查询该模具的参与人员及用时
func GetAllByMold(moldsn string) (res []Result, err error) {
    //var res []Result
    o := orm.NewOrm()
    q := "SELECT b.empno, b.empname, c.moldsn, c.moldname, d.codeno, d.descrip, "+
        "date_format(a.timestarted, '%Y-%c-%d %h:%i:%s') as timestarted, "+
        "date_format(a.timecompleted, '%Y-%c-%d %h:%i:%s') as timecompleted, "+
        "timestampdiff(minute, a.timestarted, "+
        "a.timecompleted)/60 as cal "+
        "FROM joblog a, empinfo b, moldinfo c, jobcodes d "+
        "where a.empid = b.id and a.moldid = c.id and a.jobid = d.id and "+
        "c.moldsn = ?"
    _, e := o.Raw(q, moldsn).QueryRows(&res)
    if e == nil {
        return res, nil
    }
    return nil, e
}

//通过人员编号查询其参与的项目及用时
func GetAllByEmp(empno string) (res []Result, err error) {
    //var res []Result
    o := orm.NewOrm()
    q := "SELECT b.empno, b.empname, c.moldsn, c.moldname, d.codeno, d.descrip, "+
        "date_format(a.timestarted, '%Y-%c-%d %h:%i:%s') as timestarted, "+
        "date_format(a.timecompleted, '%Y-%c-%d %h:%i:%s') as timecompleted, "+
        "timestampdiff(minute, a.timestarted, "+
        "a.timecompleted)/60 as cal "+
        "FROM joblog a, empinfo b, moldinfo c, jobcodes d "+
        "where a.empid = b.id and a.moldid = c.id and a.jobid = d.id and "+
        "b.empno = ?"
    _, e := o.Raw(q, empno).QueryRows(&res)
    if e == nil {
        return res, nil
    }
    return nil, e
}