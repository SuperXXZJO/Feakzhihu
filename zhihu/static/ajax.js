function ajax({method = "get",url,date,success,error}){
    var xhr = new XMLHttpRequest();
    if(date){
        date = querystring(date);
    }
    if(method == "get"&& date){
        url += "?" + date
    }
    xhr.open(method,url,true);

    if(method == "get"){
        xhr.send();
    }else{
        xhr.setRequestHeader("content-type","application/x-www-form-urlencoded")
        xhr.send(date);
    }
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4){
            if(xhr.status == 200){
               if(success){
                success(xhr.responseText);
            }else{
                if(error){
               error("Error:" + xhr.status);
            }
        }
    }
}
function querystring(obj){
    var str = "";
    for(var attr in obj){
        str += attr + "=" + obj[attr] + "&";

    }
    return str.substring(0,str.length - 1);
}}}