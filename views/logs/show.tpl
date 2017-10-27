<div class="w3-card-4">
  <header class="w3-container w3-green">
    <h2>选项</h2>
  </header>
  <div class="w3-container">
    <p>
    <input class="w3-radio" type="radio" name="option" id="mnordo" value="mno" onclick="f2()">
    <label>模具</label></p>
    <p>
    <input class="w3-radio" type="radio" name="option" id="enordo" value="eno" onclick="f1()">
    <label>人员</label></p>
    <p>
    <select id="search" class="w3-select" name="searchid">
    </select></p>
    <p>
    <button class="w3-btn w3-ripple w3-green" onclick="f3()">确定</button></p>
  </div>

  <div class="w3-container">
    <table id="result" class="w3-table-all w3-hoverable">

    </table>
  </div>
  &nbsp;
  <footer class="w3-container w3-green" id="displayres">
  </footer>
</div>
&nbsp;