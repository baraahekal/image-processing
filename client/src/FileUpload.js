import React, { useState, useRef } from "react";
import { FileUpload } from 'primereact/fileupload';
import { Toast } from 'primereact/toast';
import "primereact/resources/themes/lara-dark-teal/theme.css";
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import Stepper from './Stepper';

function Upload({ selectedFilter }) {
    const [originalImage, setOriginalImage] = useState(null);
    const [editedImage, setEditedImage] = useState(null);
    const [file, setFile] = useState(null); // state to store the file
    let [step, setStep] = useState(0); // state to store the current step
    const toast = useRef(null);

    const onUpload = () => {
        toast.current.show({ severity: 'info', summary: 'Success', detail: 'File Uploaded' });
    };

 const onUploadHandler = (files) => {
    const img = new Image();
    img.onload = () => {
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');

        // Set the canvas dimensions to the desired size
        canvas.width = 500; // width
        canvas.height = 500; // height

        // Draw the image on the canvas
        ctx.drawImage(img, 0, 0, canvas.width, canvas.height);

        // Get the data URL of the resized image
        const resizedImage = canvas.toDataURL();

        // Display the resized image
        setOriginalImage(resizedImage);

        // Save the file locally
        setFile(new File([dataURItoBlob(resizedImage)], files.files[0].name, { type: files.files[0].type }));

        // Reset the step
        if (step > 0) {
            setStep(0);
        }
    };
    img.src = URL.createObjectURL(files.files[0]);
};

// Helper function to convert data URI to blob
function dataURItoBlob(dataURI) {
    const byteString = atob(dataURI.split(',')[1]);
    const mimeString = dataURI.split(',')[0].split(':')[1].split(';')[0];
    const ab = new ArrayBuffer(byteString.length);
    const ia = new Uint8Array(ab);
    for (let i = 0; i < byteString.length; i++) {
        ia[i] = byteString.charCodeAt(i);
    }
    return new Blob([ab], { type: mimeString });
}

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
                console.log(step);
                // Move to the next step in the stepper
                setStep(step + 1);
                console.log(step);
            } catch (error) {
                console.error('There was a problem with the fetch operation: ' + error.message);
            }
        };
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