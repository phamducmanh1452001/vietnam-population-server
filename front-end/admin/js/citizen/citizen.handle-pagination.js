function showCitizenOfArea(_token) {
    let religions = {'0': 'Không', '1': 'Phật giáo', '2': 'Công giáo', '3': 'Hồi giáo', '4': 'Tin lành', '5': 'Cao Đài', '6': 'Tôn giáo dân gian', '7': 'Hoà hảo', '8': 'Khác'};

    function getUrl(page, limit, avatar) {
        return `https://www.phorifai.xyz/api/citizens?page=${page.toString()}&limit=${limit.toString()}`;
    }

    // show personal detail
    function showPersonalDetail(citizen) { // checked
        let showDiv = document.querySelector('#show-info-detail');
        showDiv.className = 'card shadow mb-1';
        let subShowDiv = showDiv.querySelectorAll('div')[0];
        subShowDiv.innerHTML = '';
        let spans = new Array(10);
        for (let i = 0; i < spans.length; ++i) {
            spans[i] = document.createElement('span');
            spans[i].style.display = 'block';
            spans[i].style.padding = '3px 5px';
        }
        spans[0].innerHTML = `Mã số: ${citizen['code']}`;
        spans[1].innerHTML = `Họ tên: ${citizen['first_name']} ${citizen['middle_name']} ${citizen['last_name']}`;
        spans[2].innerHTML = `Tuổi: ${citizen['age']}`;
        spans[3].innerHTML = `Giới tính: ` + (citizen['gender'] === 'M' ? 'Nam' : 'Nữ');
        spans[4].innerHTML = `Địa chỉ: ${citizen['address']}`;
        spans[5].innerHTML = `Tôn giáo: ${religions[citizen['religion_id']]}`;
        if (citizen['collaborator_name'] !== '' && citizen['collaborator_phone'] !== '') {
            spans[6].innerHTML = `Tên cộng tác viên: ${citizen['collaborator_name']}`;
            spans[7].innerHTML = `Số điện thoại: ${citizen['collaborator_phone']}`;
        }
        spans[8].innerHTML = `Địa chỉ tạm trú: ${citizen['temporary_address']}`; 
        spans[9].innerHTML = `Ảnh`;
        for (let i = 0; i < spans.length; ++i) {
            subShowDiv.appendChild(spans[i]);
        }
    }

    // set permission 
    function addEventEdit(span2, citizen) { // checking
        span2.onclick = (e) => {
            document.querySelector('#first_name').value = citizen['first_name'];
            document.querySelector('#middle_name').value = citizen['middle_name'];
            document.querySelector('#last_name').value = citizen['last_name'];
            document.querySelectorAll('#gender option')[(citizen['gender'] === 'M' ? 1 : 0)].selected = true;
            document.querySelector('#weight').value = parseInt(citizen['weight']);
            document.querySelectorAll('#religions option')[parseInt(citizen['religion_id'])] = true;
            document.querySelector('#major').value = citizen['major'];
            if (document.querySelector('#collaborator').value === '1') {
                document.querySelector('#collaborator_name').value = citizen['collaborator_name'];
                document.querySelector('#collaborator_phone').value = citizen['collaborator_phone'];
            }
            document.querySelector('#cccd').value = citizen['code'];
            document.querySelector('#btn-add-citizen').innerHTML = 'Sửa';
        }
    }

    // show personal information
    function addEventPersonalDetail(tr, citizen) {
        tr.querySelector('td:nth-child(2)').onclick = () => {
            showPersonalDetail(citizen);
        }
    }

    function addEvent(span1, code) {
        async function deleteCitizen() {
            let response = await fetch(`https://www.phorifai.xyz/api/delete-citizen?code=${code}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': 'Bearer ' + localStorage.getItem('token')
                }
            });
            let data = await response.json();
            return data;
        }

        span1.onclick = (e) => {
            deleteCitizen()
                .then(data => {
                    alert(data['message']);
                    // showCitizenOfArea(_token);
                })
        }
    }

    function showCitizen(list_citizen) { // list_cadres is an array consists of objects
        document.querySelector('tbody').innerHTML = '';
        for (let i = 0; i < list_citizen['data'].length; ++i) {
            let tr = document.createElement('tr');
            let td = new Array(7);
            for (let j = 0; j < td.length; ++j) {
                td[j] = document.createElement('td');
            }
            // console.log(list_cadres['data'][i]['name']);
            td[0].innerHTML = list_citizen['data'][i]['code'];
            td[1].innerHTML = `${list_citizen['data'][i]['first_name']} ${list_citizen['data'][i]['middle_name']} ${list_citizen['data'][i]['last_name']}`;
            td[2].innerHTML = list_citizen['data'][i]['date_of_birth'].split('-').reverse().join('-');
            td[3].innerHTML = (list_citizen['data'][i]['gender'] === 'M' ? 'Nam' : 'Nữ');
            td[4].innerHTML = list_citizen['data'][i]['date_of_joining'].split('-').reverse().join('-');
            td[5].innerHTML = list_citizen['data'][i]['major'];
            td[6].className = 'd-none';
            let span1 = document.createElement('span');
            let span2 = document.createElement('span');
            span1.innerHTML = '<i class="fas fa-minus-circle"></i>Delete';
            span2.setAttribute('data-toggle', 'modal');
            span2.setAttribute('data-target', '#myModal');
            span2.innerHTML = '<i class="far fw fa-edit"></i>Edit';
            if (parseInt(localStorage.getItem('permission')) === 0) {
                td[6].className = 'd-none';
                document.querySelector('thead tr th:nth-child(7)').className = 'd-none';
                document.querySelector('tfoot tr th:nth-child(7)').className = 'd-none'; 
            } 
            td[6].appendChild(span1);
            td[6].appendChild(span2);
            if (parseInt(localStorage.getItem('level')) === 3 && parseInt(localStorage.getItem('permission')) !== 0) {
                td[6].className = 'd-flex justify-content-between';
                document.querySelector('thead tr th:nth-child(7)').className = '';
                document.querySelector('tfoot tr th:nth-child(7)').className = '';
            }
            addEvent(span1, td[0].innerHTML);
            addEventEdit(span2, list_citizen['data'][i]);
            for (let j = 0; j < td.length; ++j) {
                tr.appendChild(td[j]);
            }
            addEventPersonalDetail(tr, list_citizen['data'][i])
            document.querySelector('tbody').appendChild(tr);
        }
    }

    function countPage(numofrecords, numofdisplays) {
        if (parseInt(numofdisplays) > parseInt(numofrecords)) {
            return 1;
        }
        if (parseInt(numofrecords) % parseInt(numofdisplays) == 0) {
            return parseInt(numofrecords) / parseInt(numofdisplays);
        }
        return parseInt(parseFloat(numofrecords) / parseFloat(numofdisplays)) + 1;
    }

    fetch(getUrl(1, 10), {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + _token
            }
        })
        .then(response => response.json())
        .then(list_citizen => {
            // tổng dân số
            document.querySelector('#total_citizen').innerHTML = `Tổng: ${list_citizen['amount']} dân số`;

            // console.log(list_cadres);
            let cntPages = countPage(list_citizen['amount'], 10);
            // console.log(cntPages);
            let option = new Array(cntPages);
            for (let i = 0; i < cntPages; ++i) {
                option[i] = document.createElement('option');
                option[i].value = (i + 1).toString();
                option[i].innerHTML = (i + 1).toString();
                document.querySelector('#page-number').appendChild(option[i]);
            }
            showCitizen(list_citizen);
        })

    let numOfPagesEle = document.querySelector('#page-number');
    let numOfDisplaysEle = document.querySelector('#data_tables');

    numOfPagesEle.onchange = (e) => {
        let optionEle = e.target;
        fetch(getUrl(optionEle.value, numOfDisplaysEle.value), {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + _token
                }
            })
            .then(response => response.json())
            .then(list_citizen => {
                showCitizen(list_citizen);
            })
    }

    numOfDisplaysEle.onchange = (e) => {
        let optionEle = e.target;
        // console.log(optionEle.value);

        fetch(getUrl(numOfPagesEle.value, optionEle.value), {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + _token
                }
            })
            .then(response => response.json())
            .then(list_citizen => {
                document.querySelector('#page-number').innerHTML = '';
                let cntPages = countPage(list_citizen['amount'], optionEle.value);
                let option = new Array(cntPages);
                for (let i = 0; i < cntPages; ++i) {
                    option[i] = document.createElement('option');
                    option[i].value = (i + 1).toString();
                    option[i].innerHTML = (i + 1).toString();
                    document.querySelector('#page-number').appendChild(option[i]);
                }
                fetch(getUrl(numOfPagesEle.value, optionEle.value), {
                        method: 'GET',
                        headers: {
                            'Authorization': 'Bearer ' + _token
                        }
                    })
                    .then(response1 => response1.json())
                    .then(list_citizen1 => {
                        document.querySelector('#page-number').innerHTML = '';
                        let cntPages1 = countPage(list_citizen1['amount'], optionEle.value);
                        let option1 = new Array(cntPages1);
                        for (let i = 0; i < cntPages1; ++i) {
                            option1[i] = document.createElement('option');
                            option1[i].value = (i + 1).toString();
                            option1[i].innerHTML = (i + 1).toString();
                            document.querySelector('#page-number').appendChild(option1[i]);
                        }
                        showCitizen(list_citizen1);
                    })
            })
    }
}
