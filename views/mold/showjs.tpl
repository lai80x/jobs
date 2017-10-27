<script src="/static/js/jquery-1.12.4.js"></script>
<script src="/static/js/jquery.dynatable.js"></script>
<script>
    $.post(
        "/mold/getall",
        function(data) {
            for (var i=0; i<data.length; i++) {
               var row = $('<tr><td>'+data[i].Id+'</td><td>'+data[i].Moldsn+'</td><td>'+data[i].Moldname+
                  '</td></tr>');
               $('#example').append(row);
            }
        },
        "json"
    );
</script>
