
"use client"

import { 
    Dialog, 
    DialogContent, 
    DialogDescription, 
    DialogFooter, 
    DialogHeader, 
    DialogTitle
} from "@/components/ui/dialog"
import { Button } from "../ui/button";
import { useCallback, useEffect, useState } from "react";
import { ProgressBar } from "../ui/progress-bar";


interface ModalProps {
    title?: string;
    description?: string;
    currentStep?: number;
    totalSteps?: number; 
    isOpen? : boolean;
    onClose?: () => void;
    onSubmit: () => void;
    body?: React.ReactElement;
    footer?: React.ReactElement;
    actionLabel?: string;
    disabled?: boolean;
    secondaryAction?: () => void;
    secondaryActionLabel?: string;
    actionLabelVariant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link';
}

export const Modal = ({
    isOpen,
    onClose,
    onSubmit,
    title,
    description,
    body,
    footer,
    actionLabel,
    currentStep,
    totalSteps,
    disabled,
    secondaryAction,
    actionLabelVariant = 'default',
    secondaryActionLabel,
}: ModalProps) => {

    const [showModal, setShowModal] = useState(isOpen);

    useEffect(() => {
        setShowModal(isOpen);
    }, [isOpen]);

    if (!isOpen) {
        return null;
    }


    return (
        <Dialog open={isOpen} onOpenChange={onClose}>
            <DialogContent>

                <DialogHeader>
                    <DialogTitle>
                        {title}
                    </DialogTitle>
                    <DialogDescription>
                        {description}
                    </DialogDescription>
                </DialogHeader>

                <div>
                    {totalSteps! > 1 ? (
                        <div className="flex flex-col gap-12 items-center">
                            <div className="w-[95%] mt-6 dark:bg-neutral-800 bg-gray-200 h-[2px] rounded-full">
                                <ProgressBar 
                                    currentStep={currentStep!} 
                                    totalSteps={totalSteps!} 
                                />
                            </div>
                            <div className="w-full">
                                {body}
                            </div>
                        </div>
                    ) : (
                        <div>
                            {body}
                        </div>
                    )}
                    
                </div>

                <DialogFooter className="gap-2">
                    {secondaryActionLabel && (
                        <Button 
                            className="w-full"
                            variant={'outline'} onClick={secondaryAction}
                            disabled={disabled}
                        >
                            {secondaryActionLabel}
                        </Button>
                    )}
                    <Button 
                        className="w-full"
                        //@ts-ignore
                        variant={actionLabelVariant} 
                        onClick={onSubmit}
                        disabled={disabled}
                    >
                        {actionLabel}
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    )
}
