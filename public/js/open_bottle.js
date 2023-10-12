function openBottle(messageUrl){
    const messagesURL = '/api/messages';
    const formData = new FormData();
    formData.append('messageUrl', messageUrl);
    return fetch(messagesURL, {
        method: 'POST',
        body: formData,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Cant obtain JSON.');
            }
            return response.text();
        })
        .catch(error => {
            console.error(error);
        });
}