

window.onscroll = function () {
    var t = document.documentElement.scrollTop || document.body.scrollTop;

    if (t < 150) {
        document.getElementsByClassName("head1")[0].style.display = "block";
        document.getElementsByClassName("head2")[0].style.display = "none";
        document.getElementsByClassName("r4")[0].style = "position:absolute"
        document.getElementsByClassName("r5")[0].style = "position:absolute"
    } else if (t > 400) {
        document.getElementsByClassName("head2")[0].style.display = "block";
        document.getElementsByClassName("head1")[0].style.display = "none";
        document.getElementsByClassName("r4")[0].style = "position:fixed;top:0"
        document.getElementsByClassName("r5")[0].style = "position:fixed;top:240px"
    } else {

        document.getElementsByClassName("head2")[0].style.display = "block";
        document.getElementsByClassName("head1")[0].style.display = "none";
        document.getElementsByClassName("r4")[0].style = "position:absolute"
        document.getElementsByClassName("r5")[0].style = "position:absolute"

    }
}
function gopersonalP(){
    window.location.href="http://localhost:8080/api/personal/:userid"
}

function showQ(){
    document.getElementsByClassName("pop")[0].style = "display:block"
}
function hiddenQ(){
    document.getElementsByClassName("pop")[0].style = "display:none"
}
function changeC(x) {
    x.style.color = "#757d8d"
}
function normalC(x) {
    x.style.color = "#8590a6"
}
var i = 1
function changeW(x) {

    if (i++ % 2 == 1) {
        x.innerHTML = '<img src="../static/img/pl.png">收起评论'
        document.getElementsByClassName("comment")[0].style = "display:block"
    } else {
        x.innerHTML = '<img src="../static/img/pl.png">评论 '
        document.getElementsByClassName("comment")[0].style = "display:none"
    }
}
function cancelL(x) {

    if (i++ % 2 == 1) {
        x.innerHTML = '<img src="../static/img/xh.png">取消喜欢'
    } else {
        x.innerHTML = '<img src="../static/img/xh.png">喜欢'
    }
}


function changeBc(x) {

    if (i++ % 2 == 1) {
        x.style = "background-color: #0084ff;color:white"

    } else {
        x.style = "background-color: #e5f2ff;color: #1c8dff"

    }
}
function shareM() {

    if (i++ % 2 == 1) {

        document.getElementsByClassName("share")[0].style = "display:block"
    } else {

        document.getElementsByClassName("share")[0].style = "display:none"
    }
}
function shareBc(x) {
    x.style = "background-color:#f6f6f6;"
}
function shareN(x) {
    x.style = "background-color:white;"
}
function moreM() {

    if (i++ % 2 == 1) {

        document.getElementsByClassName("more")[0].style = "display:"
    } else {

        document.getElementsByClassName("more")[0].style = "display:none"
    }
}
function select1() {
    document.getElementsByClassName("rec")[0].style = "display:block"
    document.getElementsByClassName("flo")[0].style = "display:none"
    document.getElementsByClassName("hot")[0].style = "display:none"
}
function select2() {
    document.getElementsByClassName("rec")[0].style = "display:none"
    document.getElementsByClassName("flo")[0].style = "display:block"
    document.getElementsByClassName("hot")[0].style = "display:none"
}
function select3() {
    document.getElementsByClassName("rec")[0].style = "display:none"
    document.getElementsByClassName("flo")[0].style = "display:none"
    document.getElementsByClassName("hot")[0].style = "display:block"
}
function fol(x) {
    if (i++ % 2 == 1) {

        x.style = "background-color:#8590a6;color:white;"
        x.innerHTML = "取消关注"
    } else {

        x.style = "background-color: white;color: #0084ff;"
        x.innerHTML = "关注问题"
    }


  



function ajax({ method = "get", url, date, success, error }) {
    var xhr = new XMLHttpRequest();
    if (date) {
        date = querystring(date);
    }
    if (method == "get" && date) {
        url += "?" + date
    }
    xhr.open(method, url, true);

    if (method == "get") {
        xhr.send();
    } else {
        xhr.setRequestHeader("content-type", "application/x-www-form-urlencoded")
        xhr.send(date);
    }
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                if (success) {
                    success(xhr.responseText);
                } else {
                    if (error) {
                        error("Error:" + xhr.status);
                    }
                }
            }
        }
        function querystring(obj) {
            var str = "";
            for (var attr in obj) {
                str += attr + "=" + obj[attr] + "&";

            }
            return str.substring(0, str.length - 1);
        }
    }
}
window.onload = function () {
    var oDiv = document.getElementsByClassName("rec")[0];
    var ag = document.getElementsByClassName("ag")[0];
    var di = document.getElementsByClassName("di")[0];
    var pl = document.getElementsByClassName("fc1")[0];
    var fx = document.getElementsByClassName("fc1")[1];
    var sc = document.getElementsByClassName("fc1")[2];
    var xh = document.getElementsByClassName("fc1")[3];
    var lm = document.getElementsByClassName("lm")[0];
    var oc = document.getElementsByClassName("oc")[0];
    var tz = document.getElementsByClassName("tz")[0];
    var flo = document.getElementsByClassName("flo")[0];
  
  

    ajax({
        
        method: "get",
        url: "http://localhost:8080/Index.html",
        success: function (result) {
            var arr = JSON.parse(result);
            for (var i = 0; i < arr.length; i++) {
                var newDiv = document.createElement("div");
                var newA = document.createElement("a");
               
                newA.href = "http://localhost:8080/api/index.html/question/" + arr[i].id;
                newA.innerHTML = arr[i].Queation;
                newA
                var newP = document.createElement("p");
                newP.innerHTML = arr[i].Detail;
                newDiv.appendChild(newA);
                newDiv.appendChild(newP);
                newDiv.appendChild(ag);
                newDiv.appendChild(di);
                newDiv.appendChild(pl);
                newDiv.appendChild(fx);
                newDiv.appendChild(sc);
                newDiv.appendChild(xh);
                newDiv.appendChild(lm)
                oDiv.appendChild(newDiv);


            }
        },
        error:function(msg){
            alert(meg);
        }
    })
    ajax({
        method: "get",
        url: "http://localhost:8080/api/Index.html/Hot",
        success: function (result) {
           var res = json.parse(result);
           for(var i = 0; i< res.length;i++){
               var oH2 = document.createElement("h2");
               var newOc = document.createElement("div");
              
               newOc.innerHTML = oc.innerHTML
               oH2.innerText = i + arr[i].Queation
               tz.appendChild("oH2");
               tz.appendChild("newOc");
           }    
        },
        error:function(msg){
            alert(msg);
        }
    })
    ajax({
        method: "get",
        url: "http://localhost:8080/api/Index.html",
        success: function (result) {
           var bam = json.parse(result);
           for(var i = 0; i< res.length;i++){
               var oH2flo = document.createElement("h2");
               var newfuntmo = document.createElement("div");
              var newpo1 = document.createElement("div");
              var newcont = document.createElement("div");
               newfuntmo.innerHTML = funtmo.innerHTML;
               oH2flo.innerText = arr[i].Queation;
               newcont.innerText= arr[i].Detail;
               newpo1.appendChild("oH2");
               newpo1.appendChild("newfuntmo");
               flo.appendChild("newpo1");
           } 
        },
        error:function(meg){
            alert(msg);
        }
    })
}
 var likeC = document.getElementsByClassName("ag")[o];
 likeC.onclik = function(){
     ajax({
         method:"post",
         url:"http://localhost:8080/api/likecountadd",
         success:function(result){
            var arr = json.parse(result);
            if (i++ % 2 == 1) {
                likeC.style = "background-color: #0084ff;color:white"
               likeC.innerHTML =  "\u25B2已赞同" +arr[i].Likecount
            } else {
                DislikeC.style = "background-color: #e5f2ff;color: #1c8dff"
                DislikeC.innerHTML =  "\u25BC赞同" +arr[i].Likecount - 1
            }
           
         },
         error:function(msg){
             alert(meg);
         }
     })
 }

 var DislikeC = document.getElementsByClassName("di")[o];
 DislikeC.onclik = function(){
     ajax({
         method:"post",
         url:"http://localhost:8080/api/dislikecountadd",
         success:function(result){
            var arr = json.parse(result);
            if (i++ % 2 == 1) {
                DislikeC.style = "background-color: #0084ff;color:white"
                DislikeC.innerHTML =  "\u25BC" +arr[i].Dislikecount
            } else {
                DislikeC.style = "background-color: #e5f2ff;color: #1c8dff"
                DislikeC.innerHTML =  "\u25BC" +arr[i].Dislikecount - 1
            }
        
           
         },
         error:function(msg){
             alert(meg);
         }
     })
 }}


 var subQ = document.getElementsByClassName("subQ")[0];
 var ediA = document.querySelectorAll("div")[4].innerHTML
 subQ.onclik = function(){
     ajax({
         method:"post",
         url:"http://localhost:8080/api/questionadd",
         date:ediA,
         success:function(result){
             alert(result)
         },
         error:function(msg){
             alert(msg);
         }
     })
 }
