(function(){
    function getRand() {
        return ((Math.floor(Math.random() * 255) + 1) -1);
    }

    function getIpPort() {
        return getRand()+"."+getRand()+"."+getRand()+"."+getRand()+":"+((Math.floor(Math.random() * 65536) + 1) -1);
    }

    function getProtocol() {
        var prot = ["socks","http"]
        var one = ((Math.floor(Math.random() * 2) + 1) -1);
        return prot[one]
    }

    for (var i = 0; i < 10; i++) {
        print("Inserting: "+getIpPort() + " - " + getProtocol())
        db.ipPort.insert({"ipPort": getIpPort(), "protocol": getProtocol()});
    }
})()
