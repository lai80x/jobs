<script src="/static/js/jquery-1.12.4.js"></script>
<script src="/static/js/jquery.dynatable.js"></script>
<script>
    $.post(
        "/emp/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<tr><td>'+data[i].Id+'</td><td>'+data[i].Empno+'</td><td>'+data[i].Empname+
                  '</td></tr>');
               $('#example').append(row);
            }
        },
        "json"
    );
</script>
