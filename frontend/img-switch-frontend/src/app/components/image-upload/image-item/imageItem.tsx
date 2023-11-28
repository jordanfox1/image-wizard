import { Chip, Indicator, Text, Button } from "@mantine/core";
import { IconDownload, IconPhoto, IconTrash } from "@tabler/icons-react";
import Image from "next/image";

interface ImageItemProps {
    image: any;
    index: number;
    errors: any;
    onImageRemove: (index: number) => void;
    convertToNewFormat: any;
    desiredFormats: string[];
    vw: number;
    setDesiredFormats: (formats: any) => void;
    loading: boolean;
}

const ImageItem = ({ image, index, errors, onImageRemove, convertToNewFormat, desiredFormats, vw, setDesiredFormats, loading }: ImageItemProps) => {
    const handleChipChange = (value, index) => {
        setDesiredFormats((prevFormats) => {
            const updatedFormats = [...prevFormats];
            updatedFormats[index] = value;
            return updatedFormats;
        });
    };

    return (
        <div key={index} className="image-item-parent">
            <div className="format-select-btn-container">
                <Chip.Group defaultValue="webp" multiple={false} value={desiredFormats[index]} onChange={(value) => handleChipChange(value, index)}>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='webp'>WEBP</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='png'>PNG</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='jpeg'>JPEG</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='gif'>GIF</Chip>
                    <Chip radius="xs" size={vw > 1023 ? 'xl' : 'sm'} className="chip" value='bmp'>BMP</Chip>
                </Chip.Group>
            </div>
            <div key={index} className="image-item">

                <figure className="image-figure">
                    {errors[index] ? (
                        <Indicator
                            label="Error"
                            position="top-start"
                            size={vw > 1023 ? 50 : 20}
                            color="red"
                        >
                            <Image
                                className="image"
                                src={image.dataURL}
                                alt="your uploaded image"
                                width={280}
                                height={160}
                                layout="responsive"
                            />
                        </Indicator>
                    ) : image.converted ? (
                        <Indicator
                            label={vw > 1023 ? "Done" : ""}
                            position="top-start"
                            size={vw > 1023 ? 50 : 20}
                            color="green"
                        >
                            <Image
                                className="image"
                                src={image.dataURL}
                                alt="your uploaded image"
                                width={280}
                                height={160}
                                layout="responsive"
                            />
                        </Indicator>
                    ) : (
                        <Image
                            className="image"
                            src={image.dataURL}
                            alt="your uploaded image"
                            width={280}
                            height={160}
                            layout="responsive"
                        />
                    )}

                    <figcaption>
                        <Text size="md">{image.file?.name} </Text>
                    </figcaption>
                </figure>

                <div className="image-item__btn-wrapper" >
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} rightSection={<IconPhoto size={14} />} onClick={() => convertToNewFormat(image.dataURL, image.file.name, desiredFormats[index], index)} loading={loading}>
                        Convert to {desiredFormats[index] || "webp"}
                    </Button>
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} rightSection={<IconTrash size={14} />} onClick={() => onImageRemove(index)} loading={loading}>Remove</Button>
                    <Button className="btn-med" size={vw > 1023 ? 'xl' : 'sm'} component="a" rightSection={<IconDownload size={14} />} href={image.dataURL} download={image.file.name} loading={loading}>
                        Download
                    </Button>
                </div>
            </div>
        </div>
    )
}

export default ImageItem;