(function () {
    let target = document.getElementById("search_word");
    target.addEventListener("keyup", (e) => {
        let word = target.value;
        console.log(word);

        if (word.length === 0){
            console.log("empty");
            document.getElementById("search_result").innerHTML = "";
            return
        }

        postData(`/products/searchApi`, {word: word})
            .then(data => {
                if(!data){
                    document.getElementById("search_result").innerHTML = "";
                    return
                }else{
                    renderHTML(data);
                }
            })
            .catch(error => console.error(error));

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

    function renderHTML(result){
        let html = "";
        let display = document.getElementById("search_result");
        result.forEach((data) => {
            html += getElement(data);
        });
        display.innerHTML = html;
    }

    function getElement(data){
        console.log(data);
        let html = `
                  <li>
                    <a class="listview__element--right-icon" href="/products/show?id=${data.id}" title="hogehoge2">
                        <div class="position-right p1em">
                            <i class="icon-chevron-right color-sub"></i>
                        </div>
                        <div class="row no-space-bottom">
                            <div class="col2">
                                <div class="thumbnail thumbnail--movies">
                                    <div class="thumbnail__figure" style="background-image: url(${data.image_url});" title="${data.title}"></div>
                                </div>
                            </div>
                            <div class="col6 push6">
                                <h3 class="text-middle text-break">
                                    <span class="color-sub">${data.title}</span>
                                </h3>
                                <p class="text-xsmall text-overflow">
                                    ${data.detail}
                                </p>
                            </div>
                        </div>
                    </a>
                </li>
        `;
        return html;
    }
})();