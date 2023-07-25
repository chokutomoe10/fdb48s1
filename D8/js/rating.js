const testimonialData = [
    {
        user: "Zenin",
        quote: "Oke",
        image: "https://media.discordapp.net/attachments/305431915178098688/614225458719883298/Capture.PNG",
        rating: 2
    },
    {
        user: "Juro",
        quote: "Boleh juga",
        image: "https://data.1freewallpapers.com/download/devil-may-cry-dante-vs-vergil.jpg",
        rating: 3
    },
    {
        user: "Fang",
        quote: "Menarik sekali",
        image: "https://www.devilmaycry.com/5se/assets/images/topics-panel-vergil.jpg",
        rating: 5
    }
]

function allTestimonial() {
    let testimonialHTML = ""

    testimonialData.forEach((card) => {
        testimonialHTML += `<div class="testimonial">
    <img src="${card.image}" />
    <p class="quote">"${card.quote}"</p>
    <p class="author">- ${card.user}</p>
    <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>`
    })

    document.getElementById("testimonials").innerHTML = testimonialHTML
}

allTestimonial()

function filterTestimonial(rating) {
    let filteredTestimonialHTML = ""

    const filteredData = testimonialData.filter((card) => {
        return card.rating === rating
    })

    filteredData.forEach((card) => {
        filteredTestimonialHTML += `<div class="testimonial">
        <img src="${card.image}" />
        <p class="quote">"${card.quote}"</p>
        <p class="author">- ${card.user}</p>
        <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
    </div>
    `
    })

    document.getElementById("testimonials").innerHTML = filteredTestimonialHTML
}