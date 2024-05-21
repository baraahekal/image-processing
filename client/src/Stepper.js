import React, { useRef, useEffect } from "react";
import { Stepper } from "primereact/stepper";
import { StepperPanel } from "primereact/stepperpanel";

function StepperComponent({ step, originalImage, editedImage }) {
  const stepperRef = useRef(null);

  // Use an effect to update the active step when the step prop changes
  useEffect(() => {
    if (step > 0) {
      stepperRef.current.nextCallback();
    } else {
      stepperRef.current.prevCallback();
    }
  }, [step]);

  return (
    <div className="card flex justify-content-center align-items-center">
      <Stepper ref={stepperRef} style={{ flexBasis: "50rem" }}>
        <StepperPanel header="Original Image">
          <div className="flex flex-column h-12rem">
            {originalImage && (
              <img className="image" src={originalImage} alt="Original" />
            )}
          </div>
        </StepperPanel>
        <StepperPanel header="Edited Image">
          <div className="flex flex-column h-12rem">
            {editedImage && (
              <img className="image" src={editedImage} alt="Edited" />
            )}
          </div>
        </StepperPanel>
      </Stepper>
    </div>
  );
}

export default StepperComponent;
