import React from "react";
import Image from "next/image";
import s from "./RecentPosts.module.css";
import { MOCK_RECENT_POST } from "./const";
import ArticleDes from "../Article/ArticleDes";

const RecentPosts = () => {
  return (
    <section>
      <h2 className={s.recentPostParagraph}>Recent blog posts</h2>
      <div className={s.recentPostContainer}>
        <div className="flex flex-col gap-8">
          <Image
            className="h-48"
            width={100}
            height={0}
            src={`${MOCK_RECENT_POST[0].img}.png`}
            alt={MOCK_RECENT_POST[0].content}
          />
          <div className="h-full">
            <ArticleDes
              title={MOCK_RECENT_POST[0].title}
              content={MOCK_RECENT_POST[0].content}
              badge={[]}
            />
          </div>
        </div>
        <div className={s.recentPostContainerCol}>
          <div className="flex h-48">
            <Image
              className="h-48"
              width={100}
              height={0}
              src={`${MOCK_RECENT_POST[0].img}.png`}
              alt={MOCK_RECENT_POST[0].content}
            />
            <div className="small-margin w-full">
              <ArticleDes
                title={MOCK_RECENT_POST[0].title}
                content={MOCK_RECENT_POST[0].content}
                badge={[]}
              />
            </div>
          </div>
          <div className="flex h-48">
            <Image
              className="h-48"
              width={100}
              height={0}
              src={`${MOCK_RECENT_POST[0].img}.png`}
              alt={MOCK_RECENT_POST[0].content}
            />
            <div className="small-margin w-full">
              <ArticleDes
                title={MOCK_RECENT_POST[0].title}
                content={MOCK_RECENT_POST[0].content}
                badge={[]}
              />
            </div>
          </div>
        </div>
        <div className="grid col-span-full grid-cols-1">
          {MOCK_RECENT_POST[2].category}
        </div>
      </div>
    </section>
  );
};

export default RecentPosts;
