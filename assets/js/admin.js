(function(){
    let delete_button = document.getElementsByClassName("delete_button");
    [].forEach.call(delete_button, (ele) => {
        ele.addEventListener("click", (e) => {

            const v = e.currentTarget;
            const product_id = v.getAttribute("value");

            if(window.confirm("id:"+product_id+ "を削除します。よろしいですか？")){
                fetch("/admin/delete", {
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

})();