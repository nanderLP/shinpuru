import { useEffect, useState } from 'react';
import { useParams } from 'react-router';
import styled, { useTheme } from 'styled-components';
import { ColorTile } from '../../../components/ColorTile';
import { useApi } from '../../../hooks/useApi';
import { Guild } from '../../../lib/shinpuru-ts/src';

type Props = {};

const StyledDiv = styled.div`
  display: flex;

  > * {
    margin: 0 1em 1em 0;
  }
`;

const Loading = () => <span>loading ...</span>;

export const GuildHome: React.FC<Props> = () => {
  const { guildid } = useParams();
  const fetch = useApi();
  const theme = useTheme();
  const [isLoading, setIsLoading] = useState(true);
  const [guild, setGuild] = useState<Guild>();
  const [reportsCount, setReportsCount] = useState<number>();

  useEffect(() => {
    _fetch().catch();
  }, [guildid]);

  const _fetch = async () => {
    setIsLoading(true);
    setGuild(await fetch((c) => c.guilds.guild(guildid!)));
    setReportsCount((await fetch((c) => c.guilds.reportsCount(guildid!))).count);
    setIsLoading(false);
  };

  return (
    <>
      {(!isLoading && (
        <StyledDiv>
          <ColorTile color={theme.blurple}>{guild?.member_count} Members</ColorTile>
          <ColorTile color={theme.red}>{reportsCount} Reports</ColorTile>
        </StyledDiv>
      )) || <Loading />}
    </>
  );
};
