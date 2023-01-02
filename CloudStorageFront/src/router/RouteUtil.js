const CookieUtil=new Object({
        getCookie: function (name) {
            let cookieList = document.cookie.split(";");
            for (let i = 0; i < cookieList.length; i++) {
                let keyValueList = cookieList[i].split("=");
                let key = keyValueList[0];
                let value = keyValueList[1];
                if (key == name) {
                    return value;
                }
            }
            return "";
        },
    //expireTime base on milliSecond
    setCookie:function(key,value,expireTime){
            let expire=";expires="+new Date(new Date().getTime()+1000*expireTime).toDateString();
            document.cookie=key+"="+value+expire;
    },
    deleteCookie:function(key,value){
            let overdueTime=";expires="+new Date(0);
            document.cookie=key+"="+value+overdueTime;
    }
    }
)
export  default  CookieUtil;