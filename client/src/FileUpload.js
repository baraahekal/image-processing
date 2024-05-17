import React, { useState, useRef } from "react";
import { FileUpload } from 'primereact/fileupload';
import { Toast } from 'primereact/toast';
import "primereact/resources/themes/lara-light-cyan/theme.css";
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import Stepper from './Stepper';

function Upload({ selectedFilter }) {
    const [originalImage, setOriginalImage] = useState(null);
    const [editedImage, setEditedImage] = useState(null);
    const [file, setFile] = useState(null); // state to store the file
    const [step, setStep] = useState(0); // state to store the current step
    const toast = useRef(null);

    const onUpload = () => {
        toast.current.show({ severity: 'info', summary: 'Success', detail: 'File Uploaded' });
    };

    const onUploadHandler = (files) => {
        // Display the original image
        setOriginalImage(URL.createObjectURL(files.files[0]));

        // Save the file locally
        setFile(files.files[0]);
    };

    const onSubmitHandler = async () => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onloadend = async () => {
            const base64data = reader.result;

            if (file === null) {
                toast.current.show({ severity: 'error', summary: 'Error', detail: 'Please upload an image' });
                return;
            }

            let contentType = 'image/jpeg';
            if (file.type === 'image/png') {
                contentType = 'image/png';
            }

            try {
                const response = await fetch('http://localhost:8080/image_processing', {
                    method: 'POST',
                    headers: {
                        'Content-Type': contentType
                    },
                    body: JSON.stringify({ data: base64data, filter: selectedFilter }) // Include the selected filter
                });

                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }

                const image = await response.blob();
                setEditedImage(URL.createObjectURL(image));
                onUpload();
                nextStep(); // Move to the next step
            } catch (error) {
                console.error('There was a problem with the fetch operation: ' + error.message);
            }
        };
    };

    const nextStep = () => {
        setStep(step + 1);
    };

    return (
        <div className="card flex justify-content-center">
            <Toast ref={toast}></Toast>
            <FileUpload mode="basic" name="demo[]" accept="image/*" maxFileSize={1000000} onSelect={onUploadHandler} />
            <button onClick={onSubmitHandler}>Submit</button>
            <Stepper step={step} originalImage={originalImage} editedImage={editedImage} />
        </div>
    )
}

export default Upload;