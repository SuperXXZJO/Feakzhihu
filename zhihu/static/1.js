
function myFunction() {
    document.getElementById("tab1").style.display="block";
    document.getElementById("tab2").style.display="none";

}


function myF() {
   
        document.getElementById("tab2").style.display="block";
        document.getElementById("tab1").style.display="none";
    }
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
    }
  
        }}
        window.onload = function(){
            var oBtn1 = document.getElementsByClassName('register1')[0];
            var oBtn2 = document.getElementsByClassName('login1')[0];
            var aInput = document.getElementsByTagName("input");
            oBtn1.onclick = function(){
                ajax({
                    method: "post",
                    url: "http://localhost:8080/signup",
                    date:{
                        Phone: aInput[2].value,
                        Password: aInput[3].value,
                    },
                    success:function(result){
                        
                            window.location.href="http://localhost:8080/Index.html"
                        },
                    
                    error: function(msg){
                       
                        alert(msg);}
                })
                    
                    
                
            } 
            oBtn2.onclick = function(){
                ajax({
                    method: "post",
                    url: "http://localhost:8080/login",
                    date:{
                        Phone: aInput[2].value,
                        Password: aInput[3].value,
                    },
                    success:function(result){
                        
                            window.location.href="http://localhost:8080/Index.html"
                        },
                    
                    error: function(msg){
                        alert(msg);
                    }
                    
                    
                })
            
        }    
        }
