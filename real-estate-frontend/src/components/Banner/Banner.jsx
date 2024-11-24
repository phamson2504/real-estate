import React from 'react'
import TypeWriterEffect from "react-typewriter-effect";
import { FaSearch } from "react-icons/fa";
import "./Banner.css"

export default function () {
    return (
        <div className="banner" >
            <div className='banner-content'>
                <div className="banner-title">
                    <h2 className="title">
                        <TypeWriterEffect
                            startDelay={1000}
                            cursorColor="#3F3D56"
                            multiText={[
                                "Discover Your Dream Home with ",
                                "Your Partner in Building a Lifetime of Memories.",
                                "Stay Informed. Stay Inspired. Stay Home.",
                                "Easy way to find a perfect property.",
                            ]}
                            multiTextDelay={1000}
                            typeSpeed={30}
                        />
                    </h2>
                    <h2 className="sub-title">
                        <span className="">
                            - ABuild Homes Estates
                        </span>
                    </h2>

                    <div className="description">
                        <p >
                            Where every brick tells a unique and captivating story. We
                            believe that a home is not just a structure; its a narrative
                            waiting to unfold.
                        </p>
                    </div>
                </div>
            </div>
            <div className="form-container">
                <div className="form-item">
                    <h3>Location</h3>
                    <select name="location" id="location">
                        <option value="">Select your city</option>
                        <option value="dhk">Dhaka, BD</option>
                        <option value="gpr">Gazipur, BD</option>
                        <option value="utr">Uttara, BD</option>
                        <option value="nw">New York, USA</option>
                        <option value="wt">Washington, USA</option>
                        <option value="prs">Paris, FR</option>
                        <option value="prg">Prague, CZK</option>
                    </select>
                </div>
                <div className="form-item">
                    <h3>Property Type</h3>
                    <select name="propertyType" id="propertyType">
                        <option value="">Select your property type</option>
                        <option value="apt">Apartment</option>
                        <option value="b">Bungalow</option>
                        <option value="c">FarmHouse</option>
                        <option value="d">Industrial</option>
                        <option value="e">Ranch</option>
                    </select>
                </div>
                <div className="form-item">
                    <h3>Price Range</h3>
                    <p>Choose Price Range</p>
                </div>
                <div className="button-container">
                    <button>
                        <FaSearch className="search-icon" />
                    </button>
                </div>
            </div>
        </div>
    )
}
