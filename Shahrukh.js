import react, {Component} from 'react'
import Car from './Car.jpg'
import React from "react";


class Shahrukh extends Component{


    render(){
        return(
        <div>
            <h1>This is Shahrukh's Component</h1>
            <img src={Car} alt={"Cyber Truck"}/>
            <div><span></span></div>
            <a href={"https://coursera.community/networking-social-discussion-5/your-best-tips-and-strategies-for-working-from-home-7923"}>Here's the Link </a>
        </div>
        )
    }
}

export default Shahrukh;