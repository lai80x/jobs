<script src="/static/js/jquery.js"></script>
<script>
    $.post(
        "/emp/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<option id='+data[i].Id+' value='+data[i].Id+'>'+data[i].Empno+'_'+data[i].Empname+
                '</option>');
               $('#empselect').append(row);
            }
        },
        "json"
    );
</script>
<script>
    $.post(
        "/mold/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<option id='+data[i].Id+' value='+data[i].Id+'>'+data[i].Moldsn+'_'+data[i].Moldname+
                '</option>');
               $('#moldselect').append(row);
            }
        },
        "json"
    );
</script>
<script>
    $.post(
        "/codes/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<option id='+data[i].Id+' value='+data[i].Id+'>'+data[i].Codeno+'_'+data[i].Descrip+
                '</option>');
               $('#codesselect').append(row);
            }
        },
        "json"
    );
</script>