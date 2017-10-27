<script src="/static/js/jquery-1.12.4.js"></script>
<script src="/static/js/jquery.dynatable.js"></script>
<script>
    function f3() {//根据用户所选的条件来列出相关信息
        var x = $('#search').val();
        var mr = document.getElementById("mnordo");//模具
        var er = document.getElementById("enordo");//人员
        var url;
        var para;
        var th = '<thead><tr class=w3-green><th>薪号</th><th>姓名</th><th>模具编号</th>'+
                '<th>模具名称</th><th>工作代码</th><th>工作名称</th><th>开始时间</th>'+
                '<th>结束时间</th><th>用时</th></tr></thead>';
        if (mr.checked) {//当模具选定时
            url = "/logs/search";
            para = {mold:x};
        } else if (er.checked) {//当人员选定时
            url = "/logs/searchemp";
            para = {emp:x};
        };
        var t = 0;
        $('#displayres').empty();//清空表格footer
        $('#result').empty();//清空表格内容
        $('#result').append(th);//添加表格表头
        $.post(
        url,
        para,
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $(
                    '<tr><td>'+data[i].Empno+
                    '</td><td>'+data[i].Empname+
                    '</td><td>'+data[i].Moldsn+
                    '</td><td>'+data[i].Moldname+
                    '</td><td>'+data[i].Codeno+
                    '</td><td>'+data[i].Descrip+
                    '</td><td>'+data[i].Timestarted+
                    '</td><td>'+data[i].Timecompleted+
                    '</td><td>'+data[i].Cal+'</td></tr>'
               );
               t = t+data[i].Cal;
               $('#result').append(row);
            };
            if (mr.checked) {
                $('#displayres').append('项目参与总人数：'+data.length+'人；总用时：'+t+'小时。');
            } else {
                $('displayres').empty();
            };
            
        },
        "json"
        );

    };

    function f1() {//列出人员相关信息
        $('#search').empty();
        $.post(
        "/emp/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<option id='+data[i].Id+' value='+data[i].Empno+'>'+data[i].Empno+'_'+data[i].Empname+
                '</option>');
               $('#search').append(row);
            }
        },
        "json"
        );
    };

    function f2() {//列出模具相关信息
        $('#search').empty();
        $.post(
        "/mold/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<option id='+data[i].Id+' value='+data[i].Moldsn+'>'+data[i].Moldsn+'_'+data[i].Moldname+
                '</option>');
               $('#search').append(row);
            }
        },
        "json"
        );
    };
</script>
