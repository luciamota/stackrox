import React, { ReactElement, ReactNode } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

const variants = {
    open: { x: 0 },
    closed: { x: '100%' },
};

const transition = {
    ease: 'easeInOut',
    transition: 2,
};

export type SidePanelAnimatedAreaProps = {
    children: ReactNode;
    isDarkMode: boolean;
    isOpen: boolean;
};

/*
 * Render an area that contains the content of a side panel.
 * Assume its parent has position relative.
 * Assume it follows its main panel sibling.
 * Assume that onClickOutside (if it is needed) wraps the children, not this component,
 * because the 1/10 width area at the left is outside the children, but inside the area of this component.
 *
 * When isOpen changes from false to true:
 * A semi-transparent gray background color covers the main panel (underlay style).
 * The side panel opens from right to left.
 */
function SidePanelAnimatedArea({
    children,
    isDarkMode,
    isOpen,
}: SidePanelAnimatedAreaProps): ReactElement {
    const bgClassName = isDarkMode ? 'bg-base-0' : 'bg-base-100';

    return (
        <AnimatePresence initial={false}>
            {isOpen && (
                <div
                    className="absolute flex h-full justify-end left-0 top-0 w-full z-10"
                    style={{ backgroundColor: 'hsla(210, 15%, 34%, 0.5)' }}
                >
                    <motion.div
                        className={`${bgClassName} border-base-400 border-l h-full rounded-tl-lg shadow-sidepanel w-full lg:w-9/10`}
                        initial="closed"
                        animate="open"
                        exit="closed"
                        variants={variants}
                        transition={transition}
                    >
                        {children}
                    </motion.div>
                </div>
            )}
        </AnimatePresence>
    );
}

export default SidePanelAnimatedArea;
