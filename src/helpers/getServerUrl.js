function getServerUrl(setter) {
    fetch("config.json")
    .then(
        (res) => res.json())
    .then((config) => {
        setter(config.buffaloServer)
        console.log("BUFFALO SERVER URL SET TO " + config.buffaloServer)
    })
    .catch(()=>{
        setter(false)
    })
}

export default getServerUrl;