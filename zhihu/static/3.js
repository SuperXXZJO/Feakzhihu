window.onscroll = function () {
    var t = document.documentElement.scrollTop || document.body.scrollTop;

    if ( t < 110) {
        document.getElementsByClassName("head1")[0].style.display = "block";
        document.getElementsByClassName("head2")[0].style.display = "none";
        
        document.getElementsByClassName("r5")[0].style = "position:absolute"
    } else if (t>1530) {
        document.getElementsByClassName("head2")[0].style.display = "block";
        document.getElementsByClassName("head1")[0].style.display = "none";
        
        document.getElementsByClassName("r5")[0].style = "position:fixed;top:100px"
    } else {
        
        document.getElementsByClassName("head2")[0].style.display = "block";
        document.getElementsByClassName("head1")[0].style.display = "none";
      
        document.getElementsByClassName("r5")[0].style = "position:absolute"

    }
}
function gopersonalP(){
    window.location.href="http://localhost:8080/api/personal/:userid"
}

function editorSW(){
    document.getElementsByClassName("l3")[0].style = "display:block"
    document.getElementsByClassName("l1")[0].style = "display:none"
}
var i = 1
function fol(x){
    if (i++ % 2 == 1) {
     
   x.style="background-color:#8590a6;color:white;"
        x.innerHTML="取消关注"
     } else {
       
        x.style="background-color:#0084ff;;color:#fff;"
        x.innerHTML="关注问题"
     }
 }
 function shareM() {

    if (i++ % 2 == 1) {
     
       document.getElementsByClassName("share")[0].style = "display:block"
    } else {
     
        document.getElementsByClassName("share")[0].style = "display:none"
    }
}
function moreM() {

    if (i++ % 2 == 1) {
     
       document.getElementsByClassName("more")[0].style = "display:block"
    } else {
     
        document.getElementsByClassName("more")[0].style = "display:none"
    }
}
function shareME() {

    if (i++ % 2 == 1) {
     
       document.getElementsByClassName("shareE")[0].style = "display:block"
    } else {
     
        document.getElementsByClassName("shareE")[0].style = "display:none"
    }
}
function moreME() {

    if (i++ % 2 == 1) {
     
       document.getElementsByClassName("moreE")[0].style = "display:block"
    } else {
     
        document.getElementsByClassName("moreE")[0].style = "display:none"
    }
}
 
function folA(x){
    if (i++ % 2 == 1) {
     
   x.style="background-color:#8590a6;color:white;"
        x.innerHTML="取消关注"
     } else {
       
        x.style="background-color:#0084ff;color:white;"
        x.innerHTML="关注"
     }
 }
 
 function shareBc(x){
    x.style="background-color:#f6f6f6;"
}
function shareN(x){
x.style="background-color:white;"
}

function changeC(x) {
    x.style.color = "#757d8d"
}
function normalC(x) {
    x.style.color = "#8590a6"
}
const changeStyle = (data) => {
    data.value? document.execCommand(data.command, false, data.value):document.execCommand(data.command, false, null)
}


function  execCommand(name, args = null) {
    document.execCommand(name, false, args);
}
function  execCommandH(name) {
    document.execCommand(name, false, "H1");
}
function  execCommandP(name) {
    document.execCommand(name, false, "P");
}

function createLink() {
    let url = window.prompt('请输入链接地址');
    if (url) this.execCommand('createLink', url);
  }
 function insertImgLink() {
    let url = window.prompt('请输入图片地址');
    if (url) this.execCommand('insertImage', url);
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
}}}


window.onload = function () {
    var head2 = document.getElementsByClassName("head2")[0];
    var Hits = document.getElementsByClassName("Hits")[0];
    var wt = document.getElementsByClassName("wt")[0];
  
    var l2 = document.getElementsByClassName("l2")[0];
   var bottom =document.getElementsByClassName("bottom")[0];
    
    ajax({
        
        method: "get",
        url: "http://localhost:8080/api/index.html/queation/:id",
        success: function (result) {
            var arr = JSON.parse(result); 
               head2.innerText = arr.queation;
               Hits.innerText = arr.Hits;
               wt .innerText= arr.Response_id;
               for(var i = 0;i< arr.length; i++){
            var oP =document.createElement("p");
            var newDiv1 = document.createElement("div");//t`
            var newDiv2 = document.createElement("div");//wt
           var newDiv3 = document.createElement("div");//bottom
           newDiv3.innerHTML = bottom.innerHTML;
           newDiv2.innerText = wt.innerText;
            oP.innerText=arr[i].Content;
            newDiv1.appendChild(newDiv2);
           newDiv1.appendChild(oP);
           newDiv1.appendChild(newDiv3);
            l2.appendChild(newDiv1);


               }
            },
        
        error:function(msg){
            alert(meg);
        }
    })
}
var sbtb = document.getElementsByClassName("sbtb");
var answer = document.querySelectorAll("div")[18].innerHTML;

sbtb.onclick = function(){
    ajax({
        method:"post",
        url:"http://localhost:8080/api/responseadd",
        date:answer,
        success:function(result){
            alert(result);
        },
        error:function(msg){
            alert(msg);
        }
    })
}