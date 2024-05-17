import React, { useRef } from "react";
import { Stepper } from 'primereact/stepper';
import "primereact/resources/themes/lara-light-cyan/theme.css";
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import { StepperPanel } from 'primereact/stepperpanel';

function StepperComponent({ step, originalImage, editedImage }) {
    const stepperRef = useRef(null);

    return (
        <div className="card flex justify-content-center">
            <Stepper ref={stepperRef} activeIndex={step} style={{ flexBasis: '50rem' }}>
                <StepperPanel header="Original Image">
                    <div className="flex flex-column h-12rem">
                        {originalImage && <img src={originalImage} alt="Original" />}
                    </div>

                </StepperPanel>
                <StepperPanel header="Edited Image">
                    <div className="flex flex-column h-12rem">
                        {editedImage && <img src={editedImage} alt="Edited" />}
                    </div>

                </StepperPanel>
            </Stepper>
        </div>
    )
}

export default StepperComponent;