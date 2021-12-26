async function getProvinces() {
    let response = await fetch('https://www.phorifai.xyz/api/provinces?page=1&limit=1000');
    let data = await response.json();
    return data;
}

async function getDistricts(province_code) {
    let response = await fetch(`https://www.phorifai.xyz/api/districts?province_code=${province_code}&page=1&limit=1000`);
    let data = await response.json();
    return data;
}

async function getWards(district_code) {
    let response = await fetch(`https://www.phorifai.xyz/api/wards?district_code=${district_code}&page=1&limit=99999`);
    let data = await response.json();
    return data;
}

