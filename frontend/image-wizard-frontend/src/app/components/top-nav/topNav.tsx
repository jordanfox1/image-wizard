import React from "react";
import "./topNav.css";
import { IconPhoto, IconWand } from '@tabler/icons-react';

export const TopNav = () => {
    return (
        <section className="top-nav">
            <div className="logo-container">
                {/* <span className='eye-icon'>👁️</span> */}
                < IconWand />
                <span className="top-nav-text">Image Wizard</span>
            </div>
        </section>
    )
};
