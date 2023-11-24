import { ImageUpload } from './components/image-upload/imageUpload'
import { TopNav } from './components/top-nav/topNav';
import PageTitle from './components/page-title/pageTitle';
export default function Home() {

  return (
    <main>
      <TopNav />
      <PageTitle />
      <ImageUpload />
    </main>
  );
}
