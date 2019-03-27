function addInfo(response) {
    var div1 = document.getElementById("name")
    var div2 = document.getElementById("email")
    var div3 = document.getElementById("address")
    var div4 = document.getElementById("phone")
    console.log(response.name)
    if (response.name == "") {
        div1.append("Nombre: Without Nombre ")   
    }else{
        div1.append("Nombre: "+response.name+"")
    }
    if (response.mail=="") {
        div2.append("Email: Without Email ")   
    }else{
        div2.append("Email: "+response.mail+"")           
    }
    if (response.address=="") {   
        div3.append("Address: Without address ")
    }else{
        div3.append("Address: "+response.address+"")
    }
    if (response.phone=="") {
        div4.append("Phone:  Without phone ")   
    }else{
        div4.append("Phone: "+response.phone+"")
    }

}