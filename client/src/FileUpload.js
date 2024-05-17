import React, { useState, useRef } from "react";
import { FileUpload } from 'primereact/fileupload';
import { Toast } from 'primereact/toast';

function Upload({ selectedFilter }) {
    const [originalImage, setOriginalImage] = useState(null);
    const [uploadedImage, setUploadedImage] = useState(null);
    const [file, setFile] = useState(null); // state to store the file
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
                setUploadedImage(URL.createObjectURL(image));
                onUpload();
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
            {originalImage && <img src={originalImage} alt="Original" />}
            {uploadedImage && <img src={uploadedImage} alt="Modified" />}
        </div>
    )

    return (
        <div>
            <Toast ref={toast}></Toast>

            <Tooltip target=".custom-choose-btn" content="Choose" position="bottom" />
            <Tooltip target=".custom-upload-btn" content="Upload" position="bottom" />
            <Tooltip target=".custom-cancel-btn" content="Clear" position="bottom" />

            <FileUpload mode="basic" name="demo[]" url="/api/upload" multiple accept="image/*" maxFileSize={1000000}
                        onSelect={onUploadHandler}
                        headerTemplate={headerTemplate} itemTemplate={itemTemplate} emptyTemplate={emptyTemplate}
                        chooseOptions={chooseOptions} uploadOptions={uploadOptions} cancelOptions={cancelOptions} />
            <button onClick={onSubmitHandler}>Submit</button>
            {originalImage && <img src={originalImage} alt="Original" />}
            {uploadedImage && <img src={uploadedImage} alt="Modified" />}
        </div>
    )
}

export default Upload;