document.querySelector('#btn-logout').addEventListener('click', (e) => {
    $.ajax({
        method: 'POST',
        url: 'https://www.phorifai.xyz/api/logout',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        success: (response) => {
            // console.log(response.token); // return a string 
            if (localStorage.getItem('token_2')) {
                $.ajax({
                    method: 'POST',
                    url: 'https://www.phorifai.xyz/api/logout',
                    headers: {
                        'Authorization': 'Bearer ' + localStorage.getItem('token_1')
                    },
                    success: (response1) => {
                        $.ajax({
                            method: 'POST',
                            url: 'https://www.phorifai.xyz/api/logout',
                            headers: {
                                'Authorization': 'Bearer ' + localStorage.getItem('token_2')
                            },
                            success: (response2) => {
                                localStorage.clear();
                                window.location = '../main/html/mainpage.html';
                            },
                            error: (er) => {
                                alert('Đăng xuất thất bại!');
                            }
                        })
                    }
                })
            }
            localStorage.clear();
            window.location = '../main/html/mainpage.html';
        },
        error: (er) => {
            alert('Đăng xuất thất bại!');
        }
    })
})