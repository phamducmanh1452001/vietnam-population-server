$(document).ready(() => {
    let searchInput = document.querySelector('#search');
    let searchBtn = document.querySelector('#btn-search');

    async function getCadres(_token, key_search) {
        let response = await fetch(`https://www.phorifai.xyz/api/lower-cadres?page=1&limit=1000&key=${key_search}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + _token
            }
        });

        let data = await response.json();
        return data;
    }

    searchBtn.onclick = () => {
        let showDiv = document.querySelector('#show-info-detail');
        showDiv.className = 'card shadow mb-1';
        let subShowDiv = showDiv.querySelectorAll('div')[0];
        subShowDiv.innerHTML = '';

        let valueInput = searchInput.value.toString();

        getCadres(localStorage.getItem('token'), valueInput)
            .then(data => {
                let cadres = data['data'][0];
                let spans = new Array(7);
                for (let i = 0; i < spans.length; ++i) {
                    spans[i] = document.createElement('span');
                    spans[i].style.display = 'block';
                    spans[i].style.padding = '3px 5px';
                }
                spans[0].innerHTML = `Mã khu vực: ${cadres['code']}`;
                spans[1].innerHTML = `Chức vụ: ${cadres['name']}`;
                spans[2].innerHTML = `Tuổi: ${cadres['age']}`;
                spans[3].innerHTML = `Số điện thoại: ${cadres['phone']}`;
                spans[4].innerHTML = `Email: ${cadres['email']}`;
                spans[5].innerHTML = 'Quyền chỉnh sửa: ' + (parseInt(cadres['permission']) === 1 ? 'Có' : 'Không');
                spans[6].innerHTML = `Khu vực quản lý: ${cadres['subdivision']['name']}`;
                // spans[6].innerHTML = 'trung';
                for (let i = 0; i < spans.length; ++i) {
                    subShowDiv.appendChild(spans[i]);
                    // console.log(spans[i])
                }

            })
    }
})