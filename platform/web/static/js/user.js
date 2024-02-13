$(document).ready(function () {
    console.log("User.js loaded");

    $('#btn-logout').on("click", function (e) {
        deleteCookie('auth-session');
        window.location.href = '/user';
    });

    function getCookie(name) {
        name = name + "=";
        var cookies = document.cookie.split(';');
        for(var i = 0; i <cookies.length; i++) {
            var cookie = cookies[i];
            while (cookie.charAt(0)==' ') {
                cookie = cookie.substring(1);
            }
            if (cookie.indexOf(name) == 0) {
                return cookie.substring(name.length,cookie.length);
            }
        }
        return "";
    }

    function setCookie(name, value, expirydays) {
        var d = new Date();
        d.setTime(d.getTime() + (expirydays*24*60*60*1000));
        var expires = "expires="+ d.toUTCString();
        document.cookie = name + "=" + value + "; " + expires;
    }

    function deleteCookie(name){
        setCookie(name,"",-1);
    }
});