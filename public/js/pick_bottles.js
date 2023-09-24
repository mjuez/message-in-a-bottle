function pick_bottles(){
    const bottlesURL = '/api/bottles';
    fetch(bottlesURL)
        .then(response => {
            if (!response.ok) {
                throw new Error('Cant obtain JSON.');
            }
            return response.json();
        })
        .then(bottles => {
            const galeryContainer = document.getElementById("galery");
            // galeryContainer.classList.add("grid-container");
            // galeryContainer.innerHTML = "";
            // document.getElementById("disclamer").classList.remove("no-visible");
            bottles.forEach(bottle => {
                const imageElement = document.createElement("img");
                imageElement.src = bottle.image_url;

                const overlayContainer = document.createElement("div");
                overlayContainer.classList.add("overlay");
                overlayContainer.style.opacity = "0";
                overlayContainer.addEventListener("click", function() {
                    document.querySelectorAll(".overlay").forEach((otherOverlayContainer) => {
                        if(otherOverlayContainer !== overlayContainer){
                            otherOverlayContainer.style.opacity = "0";
                        }
                    });
                    if(overlayContainer.style.opacity === "0"){
                        overlayContainer.style.opacity = "1";
                    }else{
                        overlayContainer.style.opacity = "0";
                    }
                });
                const textContainer = document.createElement("div");
                textContainer.classList.add("overlay-text");
                const dateContainer = document.createElement("div");
                dateContainer.classList.add("date");
                dateContainer.innerHTML = bottle.date;
                const messageContainer = document.createElement("div");
                messageContainer.innerHTML = bottle.message;
                textContainer.appendChild(dateContainer);
                textContainer.appendChild(messageContainer);
                const imageContainer = document.createElement("div");
                imageContainer.classList.add("grid-item");
                
                overlayContainer.appendChild(textContainer);
                imageContainer.appendChild(imageElement);
                imageContainer.appendChild(overlayContainer);
                galeryContainer.appendChild(imageContainer);
            });
        })
        .catch(error => {
        console.error(error);
        });
}