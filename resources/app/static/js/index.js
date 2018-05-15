let index = {
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    init: function() {
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        document.addEventListener('astilectron-ready', function() {
            index.listen();

            document.getElementById("authBtn").onclick = function() {
                document.getElementById("userHash").innerHTML = 'auth process';
                index.authUser();
            }
        });
    },
    authUser: function() {
        // Create message
        let message = {
            "name": "hash",
            "payload": {
                "email": document.getElementById("email").value,
                "hashKey": document.getElementById("password").value
            }
        };

        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            asticode.loader.hide();

            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            document.getElementById("userHiddenHash").value = message.payload.hash;

            index.hideAuth();
            index.listMedia();
        })
    },
    hideAuth: function() {
        document.getElementById("authBox").remove();
    },
    listMedia: function() {
        var request = new XMLHttpRequest();
        request.open('GET', 'http://www.updemia.com/api/v1/get?hash=' + document.getElementById("userHiddenHash").value, true);

        request.onload = function() {
          if (request.status >= 200 && request.status < 400) {
            var data = JSON.parse(request.responseText);

            document.getElementById("listMedia").innerHTML = "";

            data.forEach(function(item, i){
                let div = document.createElement("div");
                div.className = "media";
                div.innerHTML = `<a href="`+ item.url +`" target="_blank" style="background-image: url('` + item.img + `');"> </a>`;
                document.getElementById("listMedia").appendChild(div)
            });
          } else {
            asticode.notifier.error('Server error');
          }
        };

        request.onerror = function() {
          asticode.notifier.error('Error');
        };

        request.send();
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            console.log(message);
            switch (message) {
                case "refresh":
                    index.listMedia();
                    console.log("resfresh done");
                    break;
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
            }
        });
    }
};
