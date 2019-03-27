
function getData(urlsPetion){   
    var request = new XMLHttpRequest();
    request.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var response = JSON.parse(request.responseText)
            console.log(response)
            if (response.code == 200){
                window.location.href = window.location+"/dashboard"
            }else{
                alert(response.message)
            }
        }
    };    
    request.open("GET",urlsPetion,true);
    request.send(null)
}

/*This fucntion is for request post to server*/ 
function postData(data,urlsPetion,redirect){   

    var request = new XMLHttpRequest();
    
    request.open("POST",urlsPetion,false);
    request.setRequestHeader('Content-Type', 'application/json');
    request.send(data)

    if (request.readyState == 4 && request.status === 200) {
        /* Convert response in json*/
        var response = JSON.parse(request.responseText)
        console.log(response)
        if (response.code == 200){
            document.cookie = "mail="+response.payload.mail
            if (redirect !=""){
                console.log(redirect)
                window.location.href = redirect
            }else{
                return response
            }
        }else{
            alert(response.message)
        }
    }
}