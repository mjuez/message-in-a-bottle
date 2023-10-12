function pickBottles(){
    const bottlesURL = '/api/bottles';
    fetch(bottlesURL)
        .then(response => {
            if (!response.ok) {
                throw new Error('Cant obtain JSON.');
            }
            return response.json();
        })
        .then(bottles => {
            const galeryContainer = document.getElementById("gallery");
            bottles.forEach(bottle => {
                
                const imageContainer = document.createElement("div");
                
                const preloadContainer = document.createElement("div");
                preloadContainer.classList.add("preload");
                imageContainer.appendChild(preloadContainer);
                const preloadElement = document.createElement("img");
                preloadElement.src = "img/preload.gif";
                preloadContainer.appendChild(preloadElement);

                imageContainer.classList.add("grid-item");
                galeryContainer.appendChild(imageContainer);

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
                overlayContainer.appendChild(textContainer);
                
                const dateContainer = document.createElement("div");
                dateContainer.classList.add("date");
                textContainer.appendChild(dateContainer);
                
                const messageContainer = document.createElement("div");
                textContainer.appendChild(messageContainer);
                
                openBottle(bottle.messageUrl).then(message => {
                    const imageElement = document.createElement("img");
                    imageElement.src = bottle.url;

                    dateContainer.innerHTML = bottle.date;
                    messageContainer.innerHTML = message;
                    
                    imageContainer.removeChild(imageContainer.firstChild);
                    imageContainer.appendChild(imageElement);
                    imageContainer.appendChild(overlayContainer);
                });
            });
        })
        .catch(error => {
            console.error(error);
        });
}