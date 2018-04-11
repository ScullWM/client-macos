var index = {
    userHash: '',

    init: function() {
        // Wait for astilectron to be ready
        document.getElementById("auth_btn").onclick = function() {
            // index.authUser();
        }

        $('#auth_btn').css('border', '5px solid red');
    }
};
