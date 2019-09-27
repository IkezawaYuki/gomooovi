(function(){
    let delete_button = document.getElementsByClassName("delete_button");
    [].forEach.call(delete_button, (ele) => {
        ele.addEventListener("click", (e) => {

            const v = e.currentTarget;
            const product_id = v.getAttribute("value");

            if(window.confirm("id:"+product_id+ "を削除します。よろしいですか？")){
                fetch("/admin/product/delete", {
                    method: "POST",
                    mode: "cors",
                    cache: "no-cache",
                    credentials: "same-origin",
                    headers: {
                        "Content-Type": "application/json; charset=utf-8",
                    },
                    redirect: "follow",
                    referrer: "no-referrer",
                    body: JSON.stringify({product_id: product_id}),
                }).then(response => response.json())
                    .catch((e) => {
                        console.log(e);
                    });
            }

        })
    });

    let delete_button_admin_user = document.getElementsByClassName("delete_button_admin_user");
    [].forEach.call(delete_button_admin_user, (ele) => {
        ele.addEventListener("click", (e) => {

            const v = e.currentTarget;
            const user_id = v.getAttribute("value");

            if(window.confirm("id:"+user_id+ "を削除します。よろしいですか？")){
                // fetch("/admin/admin_user/delete", {
                //     method: "POST",
                //     mode: "cors",
                //     cache: "no-cache",
                //     credentials: "same-origin",
                //     headers: {
                //         "Content-Type": "application/json; charset=utf-8",
                //     },
                //     redirect: "follow",
                //     referrer: "no-referrer",
                //     body: JSON.stringify({user_id: user_id}),
                // }).then(response => response.json())
                //     .catch((e) => {
                //         console.log(e);
                //     });
            }

        })
    });

    let delete_button_user = document.getElementsByClassName("delete_button_user");
    [].forEach.call(delete_button_user, (ele) => {
        ele.addEventListener("click", (e) => {

            const v = e.currentTarget;
            const user_id = v.getAttribute("value");

            if(window.confirm("id:"+user_id+ "を削除します。よろしいですか？")){
                // fetch("/admin/user/delete", {
                //     method: "POST",
                //     mode: "cors",
                //     cache: "no-cache",
                //     credentials: "same-origin",
                //     headers: {
                //         "Content-Type": "application/json; charset=utf-8",
                //     },
                //     redirect: "follow",
                //     referrer: "no-referrer",
                //     body: JSON.stringify({user_id: user_id}),
                // }).then(response => response.json())
                //     .catch((e) => {
                //         console.log(e);
                //     });
            }

        })
    });

})();