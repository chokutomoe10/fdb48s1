let Datablog = []

function submitProject(event) {
    event.preventDefault()

    let name = document.getElementById("input-projectname").value
    let startdate = document.getElementById("start-date").value
    let enddate = document.getElementById("end-date").value
    let description = document.getElementById("input-description").value
    let image = document.getElementById("input-image").files

    if (name === "") {
        return alert('Name harus diisi!')
    } else if (startdate == "") {
        return alert('Tanggal memulai proyek harus diisi!')
    } else if (enddate == "") {
        return alert('Tanggal menyelesaikan proyek harus diisi!')
    } else if (description == "") {
        return alert('Description harus diisi!')
    } else if (image == "") {
        return alert('Masukkan gambar proyek!')
    }

    let insdate = new Date(startdate)
    let inedate = new Date(enddate)

    if (insdate > inedate) {
        return alert("Masukkan input tanggal dengan benar!");
    }

    let duration = inedate.getTime() - insdate.getTime()
    let days = duration / (1000 * 60 * 60 * 24)
    let weeks = Math.floor(days / 7)
    let months = Math.floor(weeks / 4)
    let years = Math.floor(months / 12)
    let time = ""

    if (days < 7) {
        time = days + " Hari";
    } else if (days >= 7 && weeks < 4) {
        time = weeks + " Minggu";
    } else if (weeks >= 4 && months <= 12) {
        time = months + " Bulan";
    } else {
        time = years + " Tahun";
    }
    
    const nodejs_icon = '<img src="images/nodejs.png">';
    const reactjs_icon = '<img src="images/reactjs.png">';
    const socketio_icon = '<img src="images/socket_io.png">';
    const javascript_icon = '<img src="images/javascript.svg">';

    let nodejs = document.getElementById("node-js").checked ? nodejs_icon : "";
    let socketio = document.getElementById("next-js").checked ? socketio_icon : "";
    let reactjs = document.getElementById("react-js").checked ? reactjs_icon : "";
    let javascript = document.getElementById("typescript").checked ? javascript_icon : "";

    let multitech = document.querySelectorAll(".multitech:checked");
    if (multitech.length === 0) {
    return alert("Pilih Teknologi!");
    }

    image = URL.createObjectURL(image[0]);

    let Data = {
        name,
        days,
        weeks,
        months,
        years,
        duration,
        time,
        description,
        nodejs,
        socketio,
        reactjs,
        javascript,
        image,
    }

    Datablog.push(Data)
    rendersubmitblog()

    console.log(Datablog)
}

function rendersubmitblog () {
    document.getElementById("contents").innerHTML = "";

    for (let index = 0; index < Datablog.length; index++) {
        document.getElementById("contents").innerHTML += `
            <div class="card">
                <img src="${Datablog[index].image}" class="img1">
                <div class="app">
                    <a href="blog.html"><h4>${Datablog[index].name}</h4></a>
                    <p>durasi : ${Datablog[index].time}</p>
                </div>
                <p class="note">${Datablog[index].description}</p>
                <div class="logos">
                    ${Datablog[index].nodejs}
                    ${Datablog[index].socketio}
                    ${Datablog[index].reactjs}
                    ${Datablog[index].javascript}
                </div>
                <div class="card-button">
                    <button>edit</button>
                    <button>delete</button>
                </div>
            </div>
        `;
    }
}