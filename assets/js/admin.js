(function(){
    let delete_button = document.getElementsByClassName("delete_button");
    [].forEach.call(delete_button, (ele) => {
        ele.addEventListener("click", (e) => {
            const v = e.currentTarget;
            const product_id = v.getAttribute("value");
            console.log(product_id);

            postData(`/admin/delete`, {product_id: product_id})
                .then(data => {
                    console.log(data);
                })
                .catch(error => console.log(error));
        })
    });

    function postData(url, data={}){
        return fetch(url, {
            method: "POST",
            mode: "cors",
            cache: "no-cache",
            credentials: "same-origin",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            redirect: "follow",
            referrer: "no-referrer",
            body: JSON.stringify(data),
        }).then(response => response.json());
    }

})();