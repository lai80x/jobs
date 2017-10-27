<div class="w3-card-4">

  <div class="w3-container w3-green">
    <h2>添加工时记录</h2>
  </div>

  <form id="createForm" class="w3-container" action="/logs" method="post">
    <p>
    <label>人员</label>
    <select id="empselect" class="w3-select" name="empid">
        <option value=""></option>
    </select></p>

    <p>
    <label>模具</label>
    <select id="moldselect" class="w3-select" name="moldid">
        <option value=""></option>
    </select></p>

    <p>
    <label>工作代码</label>
    <select id="codesselect" class="w3-select" name="jobid">
        <option value=""></option>
    </select></p>

    <p>
    <label>开始时间</label>
    <input class="w3-input" type="datetime-local" name="timestarted">
    </p>

    <p>
    <label>结束时间</label>
    <input class="w3-input" type="datetime-local" name="timecompleted">
    </p>

    <p>
    <label>备注</label>
    <input class="w3-input" type="text" name="note"></p>

  <p>
    <input type="submit" class="w3-btn w3-padding w3-green" style="width:120px" value="提交">
  </p>
  </form>
</div>
&nbsp;