import { Button, Text } from "@mantine/core";
import { IconUpload } from "@tabler/icons-react";

interface DropzoneProps {
    isDragging: boolean;
    dragProps: any;
    onImageUpload: () => void;
    viewportWidth: number;
}

const Dropzone: React.FC<DropzoneProps> = ({
    isDragging,
    dragProps,
    onImageUpload,
    viewportWidth,
}) => {
    return (
        <div
            className="drop-zone sticky"
            style={isDragging ? { color: "red" } : undefined}
            {...dragProps}
        >
            <Button
                size={viewportWidth > 1023 ? "xl" : "sm"}
                rightSection={<IconUpload size={14} />}
                className="btn-large"
                onClick={onImageUpload}
            >
                Select Files
            </Button>
            <Text size="md" c="#3f51b5">
                Click or drop files here
            </Text>
        </div>
    );
};

export default Dropzone;