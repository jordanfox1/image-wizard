import { Spoiler, Text } from "@mantine/core";
import "./aboutText.css";
import Link from "next/link";

export const AboutText = () => {
    return (
        <section className="spoiler-container">
            <Spoiler maxHeight={0} hideLabel="Hide website details" showLabel="Show website details" initialState={false}>
                <Text size="md">
                    Img Switch is a free tool for image file conversion. Conversions to
                    and from the following file formats are supported:
                    <Link href="https://en.wikipedia.org/wiki/WebP"> WEBP</Link>,
                    <Link href="https://en.wikipedia.org/wiki/JPEG"> JPEG</Link>,
                    <Link href="https://en.wikipedia.org/wiki/PNG"> PNG</Link>,
                    <Link href="https://en.wikipedia.org/wiki/BMP_file_format">BMP</Link>,
                    <Link href="https://en.wikipedia.org/wiki/GIF"> GIF</Link>.
                    The maximum image size supported is 100MB.
                    The maximum number of images you can convert is 100 in a 10 minute window.
                    To get started, click the button below.
                    To contact the development team for tech support, suggestions or feeback please reach out through the following channels:
                    <a href="mailto:imgswitch.dev@gmail.com?subject = Feedback&body = Message"> Email</a>, <Link href="https://www.linkedin.com/in/jordanfox183">Linkedin</Link>.
                </Text>
            </Spoiler>
            
        </section>
    );
};
