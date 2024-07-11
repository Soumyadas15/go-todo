import { motion } from 'framer-motion';
import React from 'react';

interface AnimatedDivProps {
    children: React.ReactNode;
    key: string;
}

export const AnimatedDiv: React.FC<AnimatedDivProps> = ({
    children,
    key,
}: AnimatedDivProps) => {

    return (
        <motion.div
            key={key}
            initial={{ opacity: 0, x: "-50%" }}
            animate={{ opacity: 1, x: "0%" }}
            exit={{ opacity: 0, x: "100%" }}
            transition={{ duration: 0.3, ease: "easeInOut" }}
        >
            {children}
        </motion.div>
    );
}