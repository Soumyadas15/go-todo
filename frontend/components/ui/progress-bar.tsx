"use client"

import React from 'react';
import { motion, AnimatePresence } from 'framer-motion';
import { Check } from 'lucide-react';

interface ProgressBarProps {
  currentStep: number,
  totalSteps: number,
}
export const ProgressBar = ({ 
  currentStep, 
  totalSteps 
}: ProgressBarProps) => {
  const progress = (currentStep / (totalSteps - 1)) * 100;

  return (
    <div className="relative w-full mb-6">
      <div className="w-full bg-gray-200 dark:bg-neutral-700 h-[1.5px] rounded-full overflow-hidden">
        <motion.div
          className="bg-blue-500 h-2 rounded-full"
          initial={{ width: 0 }}
          animate={{ width: `${progress}%` }}
          transition={{ duration: 0.3, ease: "easeInOut" }}
        />
      </div>
      {Array.from({ length: totalSteps }).map((_, index) => (
        <div
          key={index}
          className="absolute"
          style={{ 
            left: `${(index / (totalSteps - 1)) * 100}%`, 
            transform: 'translateX(-50%)',
            top: '50%',
            marginTop: '-12px',
          }}
        >
          <AnimatePresence mode="wait">
            {currentStep >= index ? (
              <motion.span
                key="tick"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                className="text-white text-sm flex items-center justify-center p-3 bg-blue-500 rounded-full"
                transition={{ duration: 0.1 }} 
                style={{ height: '16px', width: '16px' }}
              >
                ✓
              </motion.span>
            ) : (
              <motion.span
                key="number"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                className="text-neutral-500 text-sm flex dark:text-neutral-300 items-center justify-center bg-neutral-200 dark:bg-neutral-700 p-3 rounded-full"
                transition={{ duration: 0.1 }} 
                style={{ height: '16px', width: '16px' }} 
              >
                {index + 1}
              </motion.span>
            )}
          </AnimatePresence>
        </div>
      ))}
    </div>
  );
};