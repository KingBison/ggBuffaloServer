import React from "react";  
import "../styles/gameSelector.css"

import buffaloGraphic from "../assets/buffalo image.jpg"
import tripoleyGraphic from "../assets/tripoley graphic.png"
import mexicanGraphic from "../assets/mexican image.jpg"

const GameSelector = () => {
    return(
        <div className="game-selector-main">
            <div className="buffalo-card">
                <h1>Buffalo</h1>
                <div>
                    <img src={buffaloGraphic} alt="buffalo"/>
                </div>
            </div>
            <div className="tripoley-card">
                <h1>Tripoley</h1>
                <div>
                    <img src={tripoleyGraphic} alt="tripoley"/>
                </div>
            </div>
            <div className="mexican-train-card">
                <h1>Mexican Train</h1>
                <div>
                    <img src={mexicanGraphic} alt="mexican-train"/>
                </div>
            </div>
        </div>
    )
}

export default GameSelector;