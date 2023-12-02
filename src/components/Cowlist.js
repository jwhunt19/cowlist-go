import Cow from './Cow';

const Cowlist = ({ updateCow, cows }) => {
  return (
    <div>
      {cows.map((cow) => (
        <Cow updateCow={updateCow} cow={cow} key={cow.Id} />
      ))}
    </div>
  );
};

export default Cowlist;
