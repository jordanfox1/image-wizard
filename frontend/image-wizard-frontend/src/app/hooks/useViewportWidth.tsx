'use client';
import { useState, useEffect } from 'react';

export function useViewportWidth() {
    const [width, setWidth] = useState(0);

    useEffect(() => {
        if (typeof window !== 'undefined') {
            setWidth(window.innerWidth);

            const handleResize = () => {
                setWidth(window.innerWidth);
            }

            window.addEventListener('resize', handleResize);

            return () => {
                window.removeEventListener('resize', handleResize);
            };
        }
    }, []);

    return width;
}
