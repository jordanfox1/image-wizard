import React from "react";
import "./topNav.css";
import { IconWand } from '@tabler/icons-react';

export const TopNav = () => {
    return (
        <section className="top-nav">
            <div className="logo-container">
                < IconWand />
                <span className="top-nav-text">Image Wizard</span>
            </div>
        </section>
    )
};
