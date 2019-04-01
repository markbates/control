<%= if (err) { %>
    $("#error").html("<%= err.Error() %>");
<% } else { %>
    $("#error").html("");
<% } %>
